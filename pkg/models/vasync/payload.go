package vasync

import "github.com/ethereum/go-ethereum/common"

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
	CheckData          []byte
}
