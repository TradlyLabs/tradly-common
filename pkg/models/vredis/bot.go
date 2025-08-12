package vredis

var (
	CHANNEL_BOT = "bot"
)

type BotData struct {
	BotID   string `json:"botId"`
	OrderID string `json:"orderId"`
	AppID   string `json:"appId"`
	DexID   string `json:"dexId"`
	Payload string `json:"payload"`
}
