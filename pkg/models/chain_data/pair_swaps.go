package chain_data

import (
	"time"

	"github.com/shopspring/decimal"
)

// PairSwap represents the chain_data.pair_swaps table structure
type PairSwap struct {
	PairID             int32     `gorm:"column:pair_id;primaryKey;not null"`
	Txn1h              int64     `gorm:"column:txn_1h;not null;default:0"`
	Txn4h              int64     `gorm:"column:txn_4h;not null;default:0"`
	Txn24h             int64     `gorm:"column:txn_24h;not null;default:0"`
	Txn7d              int64     `gorm:"column:txn_7d;not null;default:0"`
	Buy1h              decimal.Decimal    `gorm:"column:buy_1h;type:numeric(38,18);not null;default:0"`
	Buy4h              decimal.Decimal    `gorm:"column:buy_4h;type:numeric(38,18);not null;default:0"`
	Buy24h             decimal.Decimal    `gorm:"column:buy_24h;type:numeric(38,18);not null;default:0"`
	Buy7d              decimal.Decimal    `gorm:"column:buy_7d;type:numeric(38,18);not null;default:0"`
	Sell1h             decimal.Decimal    `gorm:"column:sell_1h;type:numeric(38,18);not null;default:0"`
	Sell4h             decimal.Decimal    `gorm:"column:sell_4h;type:numeric(38,18);not null;default:0"`
	Sell24h            decimal.Decimal    `gorm:"column:sell_24h;type:numeric(38,18);not null;default:0"`
	Sell7d             decimal.Decimal    `gorm:"column:sell_7d;type:numeric(38,18);not null;default:0"`
	TraderAmount1h     int64     `gorm:"column:trader_amount_1h;not null;default:0"`
	TraderAmount4h     int64     `gorm:"column:trader_amount_4h;not null;default:0"`
	TraderAmount24h    int64     `gorm:"column:trader_amount_24h;not null;default:0"`
	TraderAmount7d     int64     `gorm:"column:trader_amount_7d;not null;default:0"`
	Volume1h           decimal.Decimal    `gorm:"column:volume_1h;type:numeric(38,18);not null;default:0"`
	Volume4h           decimal.Decimal    `gorm:"column:volume_4h;type:numeric(38,18);not null;default:0"`
	Volume24h          decimal.Decimal    `gorm:"column:volume_24h;type:numeric(38,18);not null;default:0"`
	Volume7d           decimal.Decimal    `gorm:"column:volume_7d;type:numeric(38,18);not null;default:0"`
	VolumeChange1h     decimal.Decimal    `gorm:"column:volume_change_1h;type:numeric(38,18);not null;default:0"`
	VolumeChange4h     decimal.Decimal    `gorm:"column:volume_change_4h;type:numeric(38,18);not null;default:0"`
	VolumeChange24h    decimal.Decimal    `gorm:"column:volume_change_24h;type:numeric(38,18);not null;default:0"`
	VolumeChange7d     decimal.Decimal    `gorm:"column:volume_change_7d;type:numeric(38,18);not null;default:0"`
	NetVolume1h        decimal.Decimal    `gorm:"column:net_volume_1h;type:numeric(38,18);not null;default:0"`
	NetVolume4h        decimal.Decimal    `gorm:"column:net_volume_4h;type:numeric(38,18);not null;default:0"`
	NetVolume24h       decimal.Decimal    `gorm:"column:net_volume_24h;type:numeric(38,18);not null;default:0"`
	NetVolume7d        decimal.Decimal    `gorm:"column:net_volume_7d;type:numeric(38,18);not null;default:0"`
	NetVolumeChange1h  decimal.Decimal    `gorm:"column:net_volume_change_1h;type:numeric(38,18);not null;default:0"`
	NetVolumeChange4h  decimal.Decimal    `gorm:"column:net_volume_change_4h;type:numeric(38,18);not null;default:0"`
	NetVolumeChange24h decimal.Decimal    `gorm:"column:net_volume_change_24h;type:numeric(38,18);not null;default:0"`
	NetVolumeChange7d  decimal.Decimal    `gorm:"column:net_volume_change_7d;type:numeric(38,18);not null;default:0"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the PairSwap struct
func (PairSwap) TableName() string {
	return "chain_data.pair_swaps"
}
