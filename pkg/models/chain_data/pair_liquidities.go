package chain_data

import (
	"time"
)

// PairLiquidity represents the chain_data.pair_liquidities table structure
type PairLiquidity struct {
	PairID             int32     `gorm:"column:pair_id;primaryKey;not null"`
	Liquidity          string    `gorm:"column:liquidity;type:numeric(38,18);not null;default:0"`
	LiquidityChange1h  string    `gorm:"column:liquidity_change_1h;type:numeric(38,18);not null;default:0"`
	LiquidityChange4h  string    `gorm:"column:liquidity_change_4h;type:numeric(38,18);not null;default:0"`
	LiquidityChange24h string    `gorm:"column:liquidity_change_24h;type:numeric(38,18);not null;default:0"`
	LiquidityChange7d  string    `gorm:"column:liquidity_change_7d;type:numeric(38,18);not null;default:0"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the PairLiquidity struct
func (PairLiquidity) TableName() string {
	return "chain_data.pair_liquidities"
}
