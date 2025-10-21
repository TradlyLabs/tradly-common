package dex

import (
	"time"
)

// Dex represents the dex.dexes table structure
type Dex struct {
	DexID       string    `gorm:"column:dex_id;"`
	URL         string    `gorm:"column:url;type:text;not null"`
	Name        string    `gorm:"column:name;type:varchar;not null"`
	Hide        bool      `gorm:"column:hide;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:now()"`
	CompanyName *string   `gorm:"column:company_name;type:varchar"`
	ChainID     *int64    `gorm:"column:chain_id"`
	Logo        string    `gorm:"column:logo;type:text;not null"`
}

// TableName sets the table name for the Dex struct
func (Dex) TableName() string {
	return "dex.dexes"
}
