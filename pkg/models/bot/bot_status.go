package bot

import (
	"time"
)

// BotStatus represents the bot.bot_status table structure
type BotStatus struct {
	StatusID  int64     `gorm:"column:status_id;primaryKey;not null"`
	BotID     int64     `gorm:"column:bot_id;not null"`
	Timestamp time.Time `gorm:"column:timestamp;not null;default:now()"`
	Status    int32     `gorm:"column:status;not null"`
	Message   *string   `gorm:"column:message;type:text"`
}

// TableName sets the table name for the BotStatus struct
func (BotStatus) TableName() string {
	return "bot.bot_status"
}
