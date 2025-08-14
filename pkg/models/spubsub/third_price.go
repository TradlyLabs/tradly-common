package spubsub

type ThirdPriceEventType int

const (
	THIRD_PRICE_EVENT_SYNC ThirdPriceEventType = iota
	THIRD_PRICE_EVENT_ADD
	THIRD_PRICE_EVENT_REMOVE
	THIRD_PRICE_EVENT_UPDATE
)

type ThirdPriceEvent struct {
	Type        ThirdPriceEventType          `json:"type"`
	PriceFeeIds map[string]ThirdPriceFeeData `json:"priceFeeIds"` // key: pyth price fee id value: symbol
}

type ThirdPriceFeeData struct {
	PriceFeeId string   `json:"priceFeeId"`
	Symbol     string   `json:"symbol"`
	Tokens     []string `json:"tokens"`
}
