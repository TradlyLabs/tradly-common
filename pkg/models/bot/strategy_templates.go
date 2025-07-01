package bot

import (
	"encoding/json"
	"time"
)

// StrategyTemplate represents the bot.strategy_templates table structure
type StrategyTemplate struct {
	TemplateID        int64           `gorm:"column:template_id;primaryKey;not null"`
	StrategyType      int32           `gorm:"column:strategy_type;not null"`
	Name              string          `gorm:"column:name;type:varchar(100);not null"`
	Description       *string         `gorm:"column:description;type:text"`
	DefaultParameters json.RawMessage `gorm:"column:default_parameters;type:jsonb;not null"`
	CreatedAt         time.Time       `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the StrategyTemplate struct
func (StrategyTemplate) TableName() string {
	return "bot.strategy_templates"
}
