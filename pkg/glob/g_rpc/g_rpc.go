package g_rpc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/TradlyLabs/tradly-common/pkg/config"
	"github.com/TradlyLabs/tradly-common/pkg/rate"
	"github.com/ethereum/go-ethereum/ethclient"
)

const defaultLimit = 100 // 100 req/s

type RPCManager struct {
	RPCMap         map[int64][]*rpcLimit
	rpcMutex       sync.RWMutex
	limiterCreator func(rate float64, b int) rate.Limiter
}

// NewRPCManager creates a new RPCManager
func NewRPCManager(limiterCreator func(limit float64, b int) rate.Limiter) *RPCManager {
	return &RPCManager{
		RPCMap:         make(map[int64][]*rpcLimit),
		limiterCreator: limiterCreator,
	}
}

type rpcLimit struct {
	Limit rate.Limiter
	URL   string
	c     *ethclient.Client
}

// GetRPC gets the RPC URL for the specified chainID
// Manage all RPCs and create a separate Limiter for each RPC according to the configuration
// When getting, find the RPC with tokens remaining and return the URL. Only return an error when all RPC tokens are used up.
func (rm *RPCManager) GetRPC(chainID int64) (*ethclient.Client, error) {
	// First try to get an RPC with tokens from the cache
	rm.rpcMutex.RLock()
	if rpcs, ok := rm.RPCMap[chainID]; ok && len(rpcs) > 0 {
		// Iterate through all RPCs to find the first one with tokens
		for _, rpc := range rpcs {
			if rpc.Limit.Allow() {
				rm.rpcMutex.RUnlock()

				// If client not initialized, create a new one
				if rpc.c == nil {
					client, err := ethclient.Dial(rpc.URL)
					if err != nil {
						return nil, fmt.Errorf("failed to dial RPC: %w", err)
					}
					rpc.c = client
				}

				return rpc.c, nil
			}
		}
	}
	rm.rpcMutex.RUnlock()

	// No cache or none have tokens, need to reinitialize
	rm.rpcMutex.Lock()
	defer rm.rpcMutex.Unlock()

	// Double check to prevent other goroutines from initializing during the write lock acquisition
	if rpcs, ok := rm.RPCMap[chainID]; ok && len(rpcs) > 0 {

		// Iterate through all RPCs again to find the first one with tokens
		for _, rpc := range rpcs {
			if rpc.Limit.Allow() {
				// If client not initialized, create a new one
				if rpc.c == nil {
					client, err := ethclient.Dial(rpc.URL)
					if err != nil {
						return nil, fmt.Errorf("failed to dial RPC: %w", err)
					}
					rpc.c = client
				}

				return rpc.c, nil
			}
		}
	}

	// Get configuration
	conf := config.C()

	// Find EVM configuration for chainID
	var rpcConfigs []config.EvmConfigRPC
	found := false

	for _, evm := range conf.Evm {
		if evm.ChainID == chainID {
			rpcConfigs = evm.RPCs
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("no evm config found for chainID %d", chainID)
	}

	if len(rpcConfigs) == 0 {
		return nil, fmt.Errorf("no rpc found for chainID %d", chainID)
	}

	// Create independent rate limiters for each RPC
	var rpcs []*rpcLimit
	for _, rpcConfig := range rpcConfigs {
		if strings.Contains(rpcConfig.URL, "ws://") {
			continue
		}
		if strings.Contains(rpcConfig.URL, "wss://") {
			continue
		}
		rateLimit := rm.limiterCreator(defaultLimit, defaultLimit)
		if rpcConfig.LimitPerSecond > 0 {
			rateLimit = rm.limiterCreator(float64(rpcConfig.LimitPerSecond), rpcConfig.LimitPerSecond)
		}
		rpcs = append(rpcs, &rpcLimit{
			Limit: rateLimit,
			URL:   rpcConfig.URL,
		})
	}

	// Save to global map
	rm.RPCMap[chainID] = rpcs

	// Iterate through all RPCs to find the first one with tokens
	for _, rpc := range rpcs {
		if rpc.Limit.Allow() {
			// If client not initialized, create a new one
			if rpc.c == nil {
				client, err := ethclient.Dial(rpc.URL)
				if err != nil {
					return nil, fmt.Errorf("failed to dial RPC: %w", err)
				}
				rpc.c = client
			}

			return rpc.c, nil
		}
	}

	// All RPCs have no tokens
	return nil, fmt.Errorf("no available rpc tokens for chainID %d", chainID)
}
