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
	TokenID           string          `json:"tokenId"`
	ChainID           json.Number     `json:"chainId"`
	TokenAddress      string          `json:"tokenAddress"`
	PairAddress       string          `json:"pairAddress"`
	BlockTimestamp    int64           `json:"blockTimestamp"`
	TransactionHash   string          `json:"transactionHash"`
	LogIndex          int64           `json:"logIndex"`
	PairTokenID       string          `json:"pairTokenId"`
	DexID             string          `json:"dexId"`
	Price             decimal.Decimal `json:"price"`
	MarketPrice       decimal.Decimal `json:"marketPrice"`
	QuoteTokenPrice   decimal.Decimal `json:"quoteTokenPrice"`
	Timestamp         int64           `json:"timestamp"`
	PairID            string          `json:"pairId"`
	BaseTokenAddress  string          `json:"baseTokenAddress"`
	QuoteTokenAddress string          `json:"quoteTokenAddress"`
}
