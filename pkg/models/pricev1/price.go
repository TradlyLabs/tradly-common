package pricev1

import (
	"context"

	"github.com/TemoreIO/temore-common/pkg/services/db"
)

type PriceFee struct {
	ID     uint   `gorm:"column:id;primaryKey"`
	Symbol string `gorm:"column:symbol"`
	FeeID  string `gorm:"column:fee_id"`
}

// TableName sets the table name for the PriceFee struct
func (PriceFee) TableName() string {
	return "pricev1.price_fees"
}

func FetchAllPriceFees(ctx context.Context) ([]*PriceFee, error) {
	var priceFees []*PriceFee
	err := db.Get().WithContext(ctx).Find(&priceFees).Error
	if err != nil {
		return nil, err
	}
	return priceFees, nil
}
