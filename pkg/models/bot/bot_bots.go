package bot

import (
	"time"
)

// BotBot represents the bot.bot_bots table structure
type BotBot struct {
	BotID         int64     `gorm:"column:bot_id;primaryKey;not null"`
	UserID        int64     `gorm:"column:user_id;not null"`
	Name          string    `gorm:"column:name;type:varchar(255);not null"`
	SourceBotID   *int64    `gorm:"column:source_bot_id"`
	StrategyType  int32     `gorm:"column:strategy_type;not null"`
	StrategyID    int64     `gorm:"column:strategy_id;not null"`
	PairID        int32     `gorm:"column:pair_id;not null"`
	ChainID       int64     `gorm:"column:chain_id;not null"`
	DexID         int32     `gorm:"column:dex_id;not null"`
	InvestToken   string    `gorm:"column:invest_token;type:char(42);not null"`
	InvestAmount  string    `gorm:"column:invest_amount;type:numeric(78,18);not null"`
	StartUSD      float64   `gorm:"column:start_usd;type:numeric(18,2);not null"`
	EndUSD        float64   `gorm:"column:end_usd;type:numeric(18,2);not null"`
	Investment    float64   `gorm:"column:investment;type:numeric(18,2);not null"`
	RealizedPNL   float64   `gorm:"column:realized_pnl;type:numeric(18,2);not null"`
	UnrealizedPNL float64   `gorm:"column:unrealized_pnl;type:numeric(18,2);not null"`
	Copiers       int32     `gorm:"column:copiers;not null"`
	Runtime       *int64    `gorm:"column:runtime"`
	Status        *int32    `gorm:"column:status;default:1"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the BotBot struct
func (BotBot) TableName() string {
	return "bot.bot_bots"
}
