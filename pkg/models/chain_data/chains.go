package chain_data

import (
	"time"
)

// Chain represents the chain_data.chains table structure
type Chain struct {
	ChainID           int64     `json:"chainId" gorm:"column:chain_id;primaryKey;not null"`
	ChainName         string    `json:"chainName" gorm:"column:chain_name;type:varchar(50);not null"`
	ChainSymbol       string    `json:"chainSymbol" gorm:"column:chain_symbol;type:varchar(10);not null"`
	RPCUrls           string    `json:"rpcUrls" gorm:"column:rpc_urls;type:text;not null"`
	LatestBlockNumber int64     `json:"latestBlockNumber" gorm:"column:latest_block_number;not null"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:created_at;not null;default:now()"`
	CurrencySymbol    string    `json:"currencySymbol" gorm:"column:currency_symbol;type:varchar(50);not null"`
	ExplorerURL       *string   `json:"explorerUrl" gorm:"column:explorer_url;type:text"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the Chain struct
func (Chain) TableName() string {
	return "chain_data.chains"
}
