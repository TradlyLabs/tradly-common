package vredis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var (
	CHANNEL_BOT = "bot"
)

type BotData struct {
	BotID             string          `json:"botId"`
	OrderID           string          `json:"orderId"`
	AppID             string          `json:"appId"`
	DexID             string          `json:"dexId"`
	Payload           string          `json:"payload"`
	PairAddress       common.Address  `json:"pairAddress"`
	RouterAddress     common.Address  `json:"routerAddress"`
	ReaderAddress     common.Address  `json:"readerAddress"`
	BaseTokenAddress  common.Address  `json:"baseTokenAddress"`
	QuoteTokenAddress common.Address  `json:"quoteTokenAddress"`
	BaseAmount        decimal.Decimal `json:"baseAmount"`
	QuoteAmount       decimal.Decimal `json:"quoteAmount"`
}
