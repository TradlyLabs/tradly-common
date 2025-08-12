package chain_data

import (
	"time"

	"github.com/shopspring/decimal"
)

// PairPrice represents the chain_data.pair_price table structure
type PairPrice struct {
	PairID            int32           `gorm:"column:pair_id;primaryKey;not null"`
	Reserve0          decimal.Decimal `gorm:"column:reserve0;type:numeric(38);not null"`
	Reserve1          decimal.Decimal `gorm:"column:reserve1;type:numeric(38);not null"`
	Price0            decimal.Decimal `gorm:"column:price0;type:numeric(38,8);not null;default:0"`
	Price1            decimal.Decimal `gorm:"column:price1;type:numeric(38,8);not null;default:0"`
	Price0Change1h    decimal.Decimal `gorm:"column:price0_change_1h;type:numeric(38,8);not null;default:0"`
	Price1Change1h    decimal.Decimal `gorm:"column:price1_change_1h;type:numeric(38,8);not null;default:0"`
	Price0Change4h    decimal.Decimal `gorm:"column:price0_change_4h;type:numeric(38,8);not null;default:0"`
	Price1Change4h    decimal.Decimal `gorm:"column:price1_change_4h;type:numeric(38,8);not null;default:0"`
	Price0Change24h   decimal.Decimal `gorm:"column:price0_change_24h;type:numeric(38,8);not null;default:0"`
	Price1Change24h   decimal.Decimal `gorm:"column:price1_change_24h;type:numeric(38,8);not null;default:0"`
	Price0Change7d    decimal.Decimal `gorm:"column:price0_change_7d;type:numeric(38,8);not null;default:0"`
	Price1Change7d    decimal.Decimal `gorm:"column:price1_change_7d;type:numeric(38,8);not null;default:0"`
	Price0Pullback1h  decimal.Decimal `gorm:"column:price0_pullback_1h;type:numeric(38,8);not null;default:0"`
	Price1Pullback1h  decimal.Decimal `gorm:"column:price1_pullback_1h;type:numeric(38,8);not null;default:0"`
	Price0Pullback4h  decimal.Decimal `gorm:"column:price0_pullback_4h;type:numeric(38,8);not null;default:0"`
	Price1Pullback4h  decimal.Decimal `gorm:"column:price1_pullback_4h;type:numeric(38,8);not null;default:0"`
	Price0Pullback24h decimal.Decimal `gorm:"column:price0_pullback_24h;type:numeric(38,8);not null;default:0"`
	Price1Pullback24h decimal.Decimal `gorm:"column:price1_pullback_24h;type:numeric(38,8);not null;default:0"`
	Price0Pullback7d  decimal.Decimal `gorm:"column:price0_pullback_7d;type:numeric(38,8);not null;default:0"`
	Price1Pullback7d  decimal.Decimal `gorm:"column:price1_pullback_7d;type:numeric(38,8);not null;default:0"`
	Price0Rebound1h   decimal.Decimal `gorm:"column:price0_rebound_1h;type:numeric(38,8);not null;default:0"`
	Price1Rebound1h   decimal.Decimal `gorm:"column:price1_rebound_1h;type:numeric(38,8);not null;default:0"`
	Price0Rebound4h   decimal.Decimal `gorm:"column:price0_rebound_4h;type:numeric(38,8);not null;default:0"`
	Price1Rebound4h   decimal.Decimal `gorm:"column:price1_rebound_4h;type:numeric(38,8);not null;default:0"`
	Price0Rebound24h  decimal.Decimal `gorm:"column:price0_rebound_24h;type:numeric(38,8);not null;default:0"`
	Price1Rebound24h  decimal.Decimal `gorm:"column:price1_rebound_24h;type:numeric(38,8);not null;default:0"`
	Price0Rebound7d   decimal.Decimal `gorm:"column:price0_rebound_7d;type:numeric(38,8);not null;default:0"`
	Price1Rebound7d   decimal.Decimal `gorm:"column:price1_rebound_7d;type:numeric(38,8);not null;default:0"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the PairPrice struct
func (PairPrice) TableName() string {
	return "chain_data.pair_price"
}
