package wallet

import "time"

// Wallet represents a wallet entity stored in database
type Wallet struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Address   string    `json:"address" gorm:"uniqueIndex"`
	Keystore  string    `json:"keystore"` // Encrypted keystore JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Wallet) TableName() string {
	return "keeper.wallets"
}

// WalletUnlock represents an unlocked wallet session
type WalletUnlock struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	WalletID  string    `json:"wallet_id" gorm:"index"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *WalletUnlock) TableName() string {
	return "keeper.wallet_unlocks"
}
