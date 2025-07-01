package authv1

import (
	"time"
)

// Wallet represents the authv1.wallets table structure
type Wallet struct {
	WalletID   int64     `gorm:"column:wallet_id;uniqueIndex;not null"`
	UserID     int64     `gorm:"column:user_id;primaryKey;not null"`
	ChainID    int64     `gorm:"column:chain_id;not null"`
	Address    string    `gorm:"column:wallet_address;type:char(42);not null"`
	Balance    string    `gorm:"column:balance;type:numeric(18,8);default:0.0"`
	WalletType *int32    `gorm:"column:wallet_type;default:1"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the Wallet struct
func (Wallet) TableName() string {
	return "authv1.wallets"
}
