package chain_data

import (
	"time"
)

// Chain represents the chain_data.chains table structure
type Chain struct {
	ChainID           int64     `gorm:"column:chain_id;primaryKey;not null"`
	ChainName         string    `gorm:"column:chain_name;type:varchar(50);not null"`
	ChainSymbol       string    `gorm:"column:chain_symbol;type:varchar(10);not null"`
	RPCUrls           string    `gorm:"column:rpc_urls;type:text;not null"`
	LatestBlockNumber int64     `gorm:"column:latest_block_number;not null"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:now()"`
	CurrencySymbol    string    `gorm:"column:currency_symbol;type:varchar(50);not null"`
	ExplorerURL       *string   `gorm:"column:explorer_url;type:text"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the Chain struct
func (Chain) TableName() string {
	return "chain_data.chains"
}
