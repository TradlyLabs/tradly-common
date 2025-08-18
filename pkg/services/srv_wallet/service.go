package srv_wallet

import (
	"context"
	"fmt"
	"os"

	"github.com/TradlyLabs/tradly-common/pkg/runtime"
	"github.com/TradlyLabs/tradly-common/pkg/wallet"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var defaultSrvWallet *SrvWallet

func init() {
	defaultSrvWallet = NewSrvWallet()
	runtime.DefaultManager.Register("SrvWallet", defaultSrvWallet)
}

type SrvWallet struct {
	walletManager *wallet.Manager
	keystoreDir   string
}

func NewSrvWallet() *SrvWallet {
	return &SrvWallet{}
}

func (s *SrvWallet) Start(ctx context.Context) error {
	dir, err := os.MkdirTemp("", "wallet-keystore")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}

	s.keystoreDir = dir

	// Create a keystore instance
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	s.walletManager = wallet.NewManager(ks)

	return nil
}

func (s *SrvWallet) Stop(ctx context.Context) error {
	if s.keystoreDir != "" {
		os.RemoveAll(s.keystoreDir)
	}

	return nil
}

func Get(args ...interface{}) *wallet.Manager {
	return defaultSrvWallet.walletManager
}
