package authv1

import (
	"time"
)

// WalletSocial represents the authv1.wallet_socials table structure
type WalletSocial struct {
	SocialID       int64     `gorm:"column:social_id;primaryKey;not null"`
	UserID         int64     `gorm:"column:user_id;not null"`
	ProviderType   int32     `gorm:"column:provider_type;not null"`
	ProviderUserID string    `gorm:"column:provider_user_id;type:varchar(255);not null"`
	Email          *string   `gorm:"column:email;type:varchar(100)"`
	Username       *string   `gorm:"column:username;type:varchar(100)"`
	AvatarURL      *string   `gorm:"column:avatar_url;type:text"`
	IsPrimary      *bool     `gorm:"column:is_primary;default:false"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the WalletSocial struct
func (WalletSocial) TableName() string {
	return "authv1.wallet_socials"
}
