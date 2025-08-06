package vredis

import "github.com/shopspring/decimal"

var (
	CHANNEL_MARKET_PAIR_PRICE = "market_pair_price"
)

// MarketPairPrice represents market data from Redis Pub/Sub
type MarketPairPrice struct {
	ID        string          `json:"id"`    // <pairID>(pair_<chainId>:<pairAddress>)
	Price     decimal.Decimal `json:"price"` // quoteToken token1 ETH in 3000 buy, 3100 sell  <-> baseToken token0 tokenIn BTC out 110000
	Timestamp int64           `json:"timestamp"`
}
