package bot

import (
	"encoding/json"
	"time"
)

// TransTransaction represents the bot.trans_transactions table structure
type TransTransaction struct {
	TransactionID int64           `gorm:"column:transaction_id;primaryKey;not null"`
	UserID        int64           `gorm:"column:user_id;not null"`
	StrategyID    *int64          `gorm:"column:strategy_id"`
	ChainID       int64           `gorm:"column:chain_id;not null"`
	DexID         int32           `gorm:"column:dex_id;not null"`
	PairID        int32           `gorm:"column:pair_id;not null"`
	PairName      string          `gorm:"column:pair_name;type:varchar(50);not null"`
	TradeType     int32           `gorm:"column:trade_type;not null"`
	Action        int32           `gorm:"column:action;not null"`
	Amount        string          `gorm:"column:amount;type:numeric(18,8);not null"`
	Price         *float64        `gorm:"column:price;type:numeric(18,8)"`
	Parameters    json.RawMessage `gorm:"column:parameters;type:jsonb;not null"`
	Status        int32           `gorm:"column:status;not null"`
	TxHash        *string         `gorm:"column:tx_hash;type:varchar(100)"`
	CreatedAt     time.Time       `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt     time.Time       `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the TransTransaction struct
func (TransTransaction) TableName() string {
	return "bot.trans_transactions"
}
