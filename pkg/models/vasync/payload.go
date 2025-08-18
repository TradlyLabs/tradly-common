package vasync

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type AutomationCompatiblePayload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	CheckData                   []byte
}

type SwapExecuteOrderPayload struct {
	ChainID       int64
	RouterAddress common.Address
	ReaderAddress common.Address
	Keeper        common.Address
	OrderID       [32]byte
	AppID         [32]byte
	DexID         [32]byte
	UserAddress   common.Address
	Nonce         *big.Int
	TokenIn       common.Address
	Amount        decimal.Decimal
	TokenOut      common.Address
	AmountOutMin  decimal.Decimal
	Exactly       bool
}

func (p *SwapExecuteOrderPayload) Valid() error {
	if p.ChainID == 0 {
		return errors.New("chain id is required")
	}
	if p.RouterAddress == (common.Address{}) {
		return errors.New("swap execute address is required")
	}
	if p.ReaderAddress == (common.Address{}) {
		return errors.New("reader address is required")
	}

	if p.Keeper == (common.Address{}) {
		return errors.New("keeper is required")
	}
	return nil
}
