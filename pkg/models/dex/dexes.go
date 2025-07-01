package dex

import (
	"time"
)

// Dex represents the dex.dexes table structure
type Dex struct {
	ID          int32     `gorm:"column:id;primaryKey;not null"`
	URL         string    `gorm:"column:url;type:text;not null"`
	Name        string    `gorm:"column:name;type:varchar;not null"`
	Hide        bool      `gorm:"column:hide;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:now()"`
	CompanyName *string   `gorm:"column:company_name;type:varchar"`
	ChainID     *int64    `gorm:"column:chain_id"`
}

// TableName sets the table name for the Dex struct
func (Dex) TableName() string {
	return "dex.dexes"
}
