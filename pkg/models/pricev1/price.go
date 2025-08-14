package pricev1

import (
	"context"

	"github.com/TradlyLabs/tradly-common/pkg/services/db"
)

type PriceFee struct {
	ID     uint                 `gorm:"column:id;primaryKey"`
	Symbol string               `gorm:"column:symbol"`
	FeeID  string               `gorm:"column:fee_id"`
	Tokens db.FlatArray[string] `gorm:"column:tokens;type:text[]"`
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

// AddTokenToPriceFee adds a token to PriceFee
func (p *PriceFee) AddTokenToPriceFee(ctx context.Context, token string) error {
	// Check if token already exists
	for _, t := range p.Tokens {
		if t == token {
			return nil // Token already exists, no need to add
		}
	}

	// Add new token
	p.Tokens = append(p.Tokens, token)

	// Save to database
	err := db.Get().WithContext(ctx).Save(p).Error
	if err != nil {
		return err
	}

	return nil
}

// RemoveTokenFromPriceFee removes a token from PriceFee
func (p *PriceFee) RemoveTokenFromPriceFee(ctx context.Context, token string) error {
	// Find and remove token
	found := false
	newTokens := []string{}
	for _, t := range p.Tokens {
		if t == token {
			found = true
			// Skip the token to be deleted
			continue
		}
		newTokens = append(newTokens, t)
	}

	// If the token to delete is not found, return directly
	if !found {
		return nil
	}
	// Update tokens
	p.Tokens = newTokens

	// Save to database
	err := db.Get().WithContext(ctx).Save(p).Error
	if err != nil {
		return err
	}

	return nil
}

// HasToken checks if PriceFee contains the specified token
func (p *PriceFee) HasToken(token string) bool {
	for _, t := range p.Tokens {
		if t == token {
			return true
		}
	}
	return false
}

// GetTokens gets all tokens
func (p *PriceFee) GetTokens() []string {
	return p.Tokens
}
