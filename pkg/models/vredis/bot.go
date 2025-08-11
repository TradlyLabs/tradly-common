package vredis

var (
	CHANNEL_BOT = "bot"
)

type BotData struct {
	BotID   string `json:"botId"`
	Payload string `json:"payload"`
}
