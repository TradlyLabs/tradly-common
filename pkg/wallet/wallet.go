// Package wallet provides wallet management functionality combining go-ethereum keystore
// with Dapr and PostgreSQL for distributed wallet management.
package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/TemoreIO/temore-common/pkg/services/db"
	"github.com/TemoreIO/temore-management/internal/models"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Manager handles wallet operations
type Manager struct {
	ks *keystore.KeyStore
	db *gorm.DB
}

// NewManager creates a new wallet manager
func NewManager(ks *keystore.KeyStore) *Manager {
	return &Manager{
		ks: ks,
		db: db.Get(),
	}
}

// CreateWallet creates a new wallet with the given password
func (m *Manager) CreateWallet(ctx context.Context, password string) (*accounts.Account, error) {
	// Create account in keystore
	account, err := m.ks.NewAccount(password)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	// Export the keystore JSON
	keyJSON, err := m.ks.Export(account, password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to export keystore: %w", err)
	}

	// Save to database
	wallet := &models.Wallet{
		ID:        uuid.New().String(),
		Address:   account.Address.Hex(),
		Keystore:  string(keyJSON),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.db.WithContext(ctx).Create(wallet).Error; err != nil {
		return nil, fmt.Errorf("failed to save wallet to database: %w", err)
	}

	return &account, nil
}

// ImportWallet imports an existing private key as a wallet
func (m *Manager) ImportWallet(ctx context.Context, privKey *ecdsa.PrivateKey, password string) (*accounts.Account, error) {
	// Import the private key into keystore
	account, err := m.ks.ImportECDSA(privKey, password)
	if err != nil {
		return nil, fmt.Errorf("failed to import private key: %w", err)
	}

	// Export the keystore JSON
	keyJSON, err := m.ks.Export(account, password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to export keystore: %w", err)
	}

	// Save to database
	wallet := &models.Wallet{
		ID:        uuid.New().String(),
		Address:   account.Address.Hex(),
		Keystore:  string(keyJSON),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.db.WithContext(ctx).Create(wallet).Error; err != nil {
		return nil, fmt.Errorf("failed to save wallet to database: %w", err)
	}

	return &account, nil
}

// GetWallet retrieves a wallet by address
func (m *Manager) GetWallet(ctx context.Context, address string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := m.db.WithContext(ctx).Where("address = ?", address).First(&wallet).Error; err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}
	return &wallet, nil
}

// ListWallets lists all wallets
func (m *Manager) ListWallets(ctx context.Context) ([]*models.Wallet, error) {
	var wallets []*models.Wallet
	if err := m.db.WithContext(ctx).Find(&wallets).Error; err != nil {
		return nil, fmt.Errorf("failed to list wallets: %w", err)
	}
	return wallets, nil
}

// DeleteWallet deletes a wallet by address
func (m *Manager) DeleteWallet(ctx context.Context, address string) error {
	// Find wallet in database
	wallet, err := m.GetWallet(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to find wallet: %w", err)
	}

	// Find account in keystore
	account := accounts.Account{
		Address: common.HexToAddress(address),
	}

	// Delete from keystore
	if err := m.ks.Delete(account, ""); err != nil {
		// Log the error but continue with database deletion
		fmt.Printf("Warning: failed to delete from keystore: %v\n", err)
	}

	// Delete from database
	if err := m.db.WithContext(ctx).Delete(&models.Wallet{}, "id = ?", wallet.ID).Error; err != nil {
		return fmt.Errorf("failed to delete wallet from database: %w", err)
	}

	return nil
}

// UnlockWallet unlocks a wallet for a specified duration
func (m *Manager) UnlockWallet(ctx context.Context, address, password string, duration time.Duration) error {
	// Find wallet in database
	wallet, err := m.GetWallet(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to find wallet: %w", err)
	}

	// Import the keystore to keystore manager
	account, err := m.ks.Import([]byte(wallet.Keystore), password, password)
	if err != nil {
		return fmt.Errorf("failed to import keystore: %w", err)
	}

	// Unlock the account
	if duration > 0 {
		err = m.ks.TimedUnlock(account, password, duration)
	} else {
		err = m.ks.Unlock(account, password)
	}

	if err != nil {
		return fmt.Errorf("failed to unlock wallet: %w", err)
	}

	// Save unlock session to database if duration > 0
	if duration > 0 {
		unlock := &models.WalletUnlock{
			ID:        uuid.New().String(),
			WalletID:  wallet.ID,
			ExpiresAt: time.Now().Add(duration),
			CreatedAt: time.Now(),
		}

		if err := m.db.WithContext(ctx).Create(unlock).Error; err != nil {
			return fmt.Errorf("failed to save unlock session: %w", err)
		}
	}

	return nil
}

// LockWallet locks a wallet
func (m *Manager) LockWallet(ctx context.Context, address string) error {
	addr := common.HexToAddress(address)

	// Lock in keystore
	if err := m.ks.Lock(addr); err != nil {
		return fmt.Errorf("failed to lock wallet: %w", err)
	}

	// Remove unlock sessions from database
	if err := m.db.WithContext(ctx).Where("wallet_id = (SELECT id FROM wallets WHERE address = ?)", address).
		Delete(&models.WalletUnlock{}).Error; err != nil {
		return fmt.Errorf("failed to remove unlock sessions: %w", err)
	}

	return nil
}

// SignHash signs a hash with the specified wallet
func (m *Manager) SignHash(ctx context.Context, address string, hash []byte) ([]byte, error) {
	addr := common.HexToAddress(address)
	account := accounts.Account{Address: addr}

	signature, err := m.ks.SignHash(account, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %w", err)
	}

	return signature, nil
}

// SignHashWithPassphrase signs a hash with the specified wallet and passphrase
func (m *Manager) SignHashWithPassphrase(ctx context.Context, address, passphrase string, hash []byte) ([]byte, error) {
	addr := common.HexToAddress(address)
	account := accounts.Account{Address: addr}

	signature, err := m.ks.SignHashWithPassphrase(account, passphrase, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %w", err)
	}

	return signature, nil
}

// GetAddresses returns all wallet addresses
func (m *Manager) GetAddresses(ctx context.Context) ([]common.Address, error) {
	// Get from database
	wallets, err := m.ListWallets(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list wallets: %w", err)
	}

	var addresses []common.Address
	for _, wallet := range wallets {
		addresses = append(addresses, common.HexToAddress(wallet.Address))
	}

	return addresses, nil
}

// IsUnlocked checks if a wallet is currently unlocked
func (m *Manager) IsUnlocked(ctx context.Context, address string) bool {
	addr := common.HexToAddress(address)
	
	// Try to find the account in keystore
	account, err := m.ks.Find(accounts.Account{Address: addr})
	if err != nil {
		return false
	}

	// Try to sign an empty hash to check if the account is unlocked
	// If the account is locked, SignHash will return ErrLocked
	_, err = m.ks.SignHash(account, []byte{})
	return err != keystore.ErrLocked
}
