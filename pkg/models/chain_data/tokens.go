package chain_data

import (
	"math/big"
	"time"
)

// Token represents the chain_data.tokens table structure
type Token struct {
	ID         string     `gorm:"column:id;primaryKey;not null"`
	Name       string     `gorm:"column:name;type:varchar(50);not null"`
	Symbol     string     `gorm:"column:symbol;type:varchar(10);not null"`
	Address    string     `gorm:"column:address;type:char(42);not null"`
	ChainID    int64      `gorm:"column:chain_id;not null"`
	CurrentUSD *big.Float `gorm:"column:current_usd;type:numeric(38,18);not null"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the Token struct
func (Token) TableName() string {
	return "chain_data.tokens"
}
