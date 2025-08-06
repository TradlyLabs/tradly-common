package spubsub

import "encoding/json"

// BotData represents the bot data to be published to pub/sub
type BotData struct {
	BotID   string `json:"botId"`
	Payload string `json:"payload"`
}

// BotDataRaw represents the bot data received from pub/sub
type BotDataRaw struct {
	BotID   string          `json:"botId"`
	Payload json.RawMessage `json:"payload"`
}
