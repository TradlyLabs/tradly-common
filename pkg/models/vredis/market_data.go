package vredis

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

var (
	CHANNEL_TOKEN_USD_PRICES = "token_usd_prices_channel"
)

// TokenUSDPrice represents market data from Redis Pub/Sub
type TokenUSDPrice struct {
	TokenID         string          `json:"tokenId"`
	ChainID         json.Number     `json:"chainId"`
	TokenAddress    string          `json:"tokenAddress"`
	PairAddress     string          `json:"pairAddress"`
	USDPrice        decimal.Decimal `json:"usdPrice"`
	BlockTimestamp  int64           `json:"blockTimestamp"`
	TransactionHash string          `json:"transactionHash"`
	LogIndex        int64           `json:"logIndex"`
}
