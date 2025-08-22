package chain_data

import (
	"context"
	"time"

	"github.com/TradlyLabs/tradly-common/pkg/services/db"
	"github.com/shopspring/decimal"
)

// Token represents the chain_data.tokens table structure
type Token struct {
	ID                string          `gorm:"column:id;primaryKey;not null"`
	Name              string          `gorm:"column:name;type:varchar(50);not null"`
	Symbol            string          `gorm:"column:symbol;type:varchar(10);not null"`
	Address           string          `gorm:"column:address;type:char(42);not null"`
	Decimals          int32           `gorm:"column:decimals;not null"`
	ChainID           int64           `gorm:"column:chain_id;not null"`
	CurrentUSD        decimal.Decimal `gorm:"column:current_usd;type:numeric(256,18);not null"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;not null;default:now()"`
	URL               string          `gorm:"column:url;type:text;not null;default:''"`
	Extensions        string          `gorm:"column:extensions;type:text;not null;default:''"`
	IsMark            bool            `gorm:"column:is_mark;not null;default:false"`
	RecentListingTime time.Time       `gorm:"column:recent_listing_time;not null;"`
	TotalSupply       decimal.Decimal `gorm:"column:total_supply;type:numeric(256,18);not null"`
	MaxSupply         decimal.Decimal `gorm:"column:max_supply;type:numeric(256,18);not null"`
}

// TableName sets the table name for the Token struct
func (Token) TableName() string {
	return "chain_data.tokens"
}

func FetchAllTokensByIDs(ctx context.Context, tokenIDs []string) ([]*Token, error) {
	var tokens []*Token
	if err := db.Get().WithContext(ctx).Where("id in ?", tokenIDs).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}
