package chain_data

import (
	"context"
	"time"

	"github.com/TemoreIO/temore-common/pkg/services/db"
)

// Pair represents the chain_data.pairs table structure
type Pair struct {
	ID            uint       `gorm:"column:id;primaryKey"`
	ChainID       int64      `gorm:"column:chain_id;not null"`
	DexID         *int32     `gorm:"column:dex_id"`
	PairAddress   string     `gorm:"column:pair_address;type:char(42);not null"`
	Symbol        string     `gorm:"column:symbol;type:char(10);not null"`
	Name          string     `gorm:"column:name;type:varchar(50);not null"`
	Factory       string     `gorm:"column:factory;type:char(42);not null"`
	Decimals      int32      `gorm:"column:decimals;not null"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null;default:now()"`
	IndexingFrom  *time.Time `gorm:"column:indexing_from"`
	Token0Address string     `gorm:"column:token0_address;type:char(42);not null"`
	Token1Address string     `gorm:"column:token1_address;type:char(42);not null"`
	Token0Symbol  string     `gorm:"column:token0_symbol;type:varchar(10);not null;default:''"`
	Token1Symbol  string     `gorm:"column:token1_symbol;type:varchar(10);not null;default:''"`
	Token0URL     string     `gorm:"column:token0_url;type:text;not null;default:''"`
	Token1URL     string     `gorm:"column:token1_url;type:text;not null;default:''"`
	Token0Name    string     `gorm:"column:token0_name;type:text;not null;default:''"`
	Token1Name    string     `gorm:"column:token1_name;type:text;not null;default:''"`
}

// TableName sets the table name for the Pair struct
func (Pair) TableName() string {
	return "chain_data.pairs"
}

func FetchAllPairs(ctx context.Context) (data []*Pair, err error) {
	err = db.Get().WithContext(ctx).Model(&Pair{}).Find(&data).Error
	return
}
