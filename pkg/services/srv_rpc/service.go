package srv_rpc

import (
	"context"

	"github.com/TradlyLabs/tradly-common/pkg/glob/g_rpc"
	"github.com/TradlyLabs/tradly-common/pkg/rate"
	"github.com/TradlyLabs/tradly-common/pkg/runtime"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RPCManager interface {
	GetRPC(chainID int64) (*ethclient.Client, error)
}

type SrvRPC struct {
	rpcManager RPCManager
}

func init() {
	runtime.DefaultManager.Register("SrvRPC", NewSrvRPC())
}

func NewSrvRPC() *SrvRPC {
	return &SrvRPC{}
}

func (s *SrvRPC) Start(context.Context) error {
	s.rpcManager = g_rpc.NewRPCManager(func(limit float64, b int) rate.Limiter {
		return rate.NewRedisLimiter("web3_rpc", limit, float64(b))

	})
	return nil
}

func (s *SrvRPC) Stop(context.Context) error {
	return nil
}

func (s *SrvRPC) GetRPC(chainID int64) (*ethclient.Client, error) {
	return s.rpcManager.GetRPC(chainID)
}
