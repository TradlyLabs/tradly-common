package spubsub

type ThirdPriceEventType int

const (
	THIRD_PRICE_EVENT_SYNC ThirdPriceEventType = iota
	THIRD_PRICE_EVENT_ADD
	THIRD_PRICE_EVENT_REMOVE
)

type ThirdPriceEvent struct {
	Type        ThirdPriceEventType `json:"type"`
	PriceFeeIds map[string]string   `json:"priceFeeIds"` // key: symbol value: pyth price fee id
}
