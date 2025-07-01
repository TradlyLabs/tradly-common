package bot

import (
	"time"
)

// TransLog represents the bot.trans_logs table structure
type TransLog struct {
	LogID         int64     `gorm:"column:log_id;primaryKey;not null"`
	TransactionID int64     `gorm:"column:transaction_id;not null"`
	Status        string    `gorm:"column:status;type:varchar(20);not null"`
	Message       *string   `gorm:"column:message;type:text"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()"`
}

// TableName sets the table name for the TransLog struct
func (TransLog) TableName() string {
	return "bot.trans_logs"
}
