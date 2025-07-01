package bot

import (
	"encoding/json"
	"time"
)

// StrategyStrategy represents the bot.strategy_strategies table structure
type StrategyStrategy struct {
	StrategyID   int64           `gorm:"column:strategy_id;primaryKey;not null"`
	StrategyType int32           `gorm:"column:strategy_type;not null"`
	UserID       int64           `gorm:"column:user_id;not null"`
	BotID        *int64          `gorm:"column:bot_id"`
	TemplateID   int64           `gorm:"column:template_id;not null"`
	Parameters   json.RawMessage `gorm:"column:parameters;type:jsonb;not null"`
	CreatedAt    time.Time       `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt    time.Time       `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the StrategyStrategy struct
func (StrategyStrategy) TableName() string {
	return "bot.strategy_strategies"
}
