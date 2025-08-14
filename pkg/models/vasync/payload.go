package vasync

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

type AutomationCompatiblePayload struct {
	NetworkName                 string
	AutomationCompatibleAddress common.Address
	Keeper                      common.Address
	CheckData                   []byte
}

type SwapExecuteOrderPayload struct {
	NetworkName        string
	SwapExecuteAddress common.Address
	Keeper             common.Address
}

func (p *SwapExecuteOrderPayload) Valid() error {
	if p.NetworkName == "" {
		return errors.New("network name is required")
	}
	if p.SwapExecuteAddress == (common.Address{}) {
		return errors.New("swap execute address is required")
	}
	if p.Keeper == (common.Address{}) {
		return errors.New("keeper is required")
	}

	return nil

}
