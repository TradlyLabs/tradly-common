package bot

import (
	"time"
)

// BotNotification represents the bot.bot_notifications table structure
type BotNotification struct {
	NotificationID   int64     `gorm:"column:notification_id;primaryKey;not null"`
	UserID           int64     `gorm:"column:user_id;not null"`
	BotID            *int64    `gorm:"column:bot_id"`
	Timestamp        time.Time `gorm:"column:timestamp;not null;default:now()"`
	Message          string    `gorm:"column:message;type:text;not null"`
	IsRead           *bool     `gorm:"column:is_read;default:false"`
	NotificationType int32     `gorm:"column:notification_type;not null"`
}

// TableName sets the table name for the BotNotification struct
func (BotNotification) TableName() string {
	return "bot.bot_notifications"
}
