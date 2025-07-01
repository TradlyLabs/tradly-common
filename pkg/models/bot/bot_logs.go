package bot

import (
	"time"
)

// BotLog represents the bot.bot_logs table structure
type BotLog struct {
	LogID     int64     `gorm:"column:log_id;primaryKey;not null"`
	BotID     int64     `gorm:"column:bot_id;not null"`
	Timestamp time.Time `gorm:"column:timestamp;not null;default:now()"`
	LogLevel  int32     `gorm:"column:log_level;not null"`
	Message   string    `gorm:"column:message;type:text;not null"`
}

// TableName sets the table name for the BotLog struct
func (BotLog) TableName() string {
	return "bot.bot_logs"
}
