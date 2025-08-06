package chain_data

import (
	"context"
	"time"

	"github.com/TradlyLabs/tradly-common/pkg/services/db"
)

// Pair represents the chain_data.pairs table structure
type Pair struct {
	ID               string     `json:"id" gorm:"column:id;primaryKey"`
	ChainID          int64      `json:"chainID" gorm:"column:chain_id;not null"`
	DexID            *int32     `json:"dexID,omitempty" gorm:"column:dex_id"`
	PairAddress      string     `json:"pairAddress" gorm:"column:pair_address;type:char(42);not null"`
	Symbol           string     `json:"symbol" gorm:"column:symbol;type:char(10);not null"`
	Name             string     `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Factory          string     `json:"factory" gorm:"column:factory;type:char(42);not null"`
	Decimals         int32      `json:"decimals" gorm:"column:decimals;not null"`
	CreatedAt        time.Time  `json:"createdAt" gorm:"column:created_at;not null;default:now()"`
	UpdatedAt        time.Time  `json:"updatedAt" gorm:"column:updated_at;not null;default:now()"`
	IndexingFrom     *time.Time `json:"indexingFrom" gorm:"column:indexing_from"`
	BaseTokenAddress string     `json:"baseTokenAddress" gorm:"column:base_token_address;type:char(42);not null"`
	Token0Address    string     `json:"token0Address" gorm:"column:token0_address;type:char(42);not null"`
	Token1Address    string     `json:"token1Address" gorm:"column:token1_address;type:char(42);not null"`
	Token0Symbol     string     `json:"token0Symbol" gorm:"column:token0_symbol;type:varchar(10);not null;default:''"`
	Token1Symbol     string     `json:"token1Symbol" gorm:"column:token1_symbol;type:varchar(10);not null;default:''"`
	Token0URL        string     `json:"token0URL" gorm:"column:token0_url;type:text;not null;default:''"`
	Token1URL        string     `json:"token1URL" gorm:"column:token1_url;type:text;not null;default:''"`
	Token0Name       string     `json:"token0Name" gorm:"column:token0_name;type:text;not null;default:''"`
	Token1Name       string     `json:"token1Name" gorm:"column:token1_name;type:text;not null;default:''"`
	Disabled         bool       `json:"disabled" gorm:"column:disabled;not null;default:false"`
}

// TableName sets the table name for the Pair struct
func (Pair) TableName() string {
	return "chain_data.pairs"
}

func FetchAllPairs(ctx context.Context) (data []*Pair, err error) {
	err = db.Get().WithContext(ctx).Model(&Pair{}).Find(&data).Error
	return
}
