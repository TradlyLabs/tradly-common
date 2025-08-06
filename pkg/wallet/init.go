// Package wallet provides wallet management functionality combining go-ethereum keystore
// with Dapr and PostgreSQL for distributed wallet management.
package wallet

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

// WalletManager is a global instance of the wallet manager
var WalletManager *Manager

// KeyStoreDir is the directory where keystore files are temporarily stored
var KeyStoreDir string

// Init initializes the wallet manager
func Init() error {
	// Create a temporary directory for keystore files
	dir, err := os.MkdirTemp("", "wallet-keystore")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}

	KeyStoreDir = dir

	// Create a keystore instance
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Create wallet manager
	WalletManager = NewManager(ks)

	return nil
}

// Cleanup cleans up temporary files
func Cleanup() error {
	if KeyStoreDir != "" {
		return os.RemoveAll(KeyStoreDir)
	}
	return nil
}

// GeneratePrivateKey generates a new private key
func GeneratePrivateKey() (*[32]byte, error) {
	key := &[32]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		return nil, err
	}

	// Ensure it's a valid private key
	for common.BytesToHash(key[:]) == (common.Hash{}) || key[0] == 0 {
		_, err = rand.Read(key[:])
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}
