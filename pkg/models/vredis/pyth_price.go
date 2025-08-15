package vredis

import "encoding/json"

func UnmarshalPrice(data []byte) (Price, error) {
	var r Price
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Price) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// A geographical coordinate
type Price struct {
	Expo int64 `json:"expo"`
	// pyth fee id
	ID string `json:"id"`
	// price with expo
	Price       string `json:"price"`
	PublishTime int64  `json:"publishTime"`
	// price symbol ex. BTC/USD, ETH/USD
	Symbol string `json:"symbol"`
}

// Pair schema start
// A geographical coordinate
func UnmarshalPair(data []byte) (Pair, error) {
	var r Pair
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Pair) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// A geographical coordinate
type Pair struct {
	// Pair address
	Address   string    `json:"address"`
	Base      *Base     `json:"base,omitempty"`
	BaseToken BaseToken `json:"baseToken"`
	Block     *Block    `json:"block,omitempty"`
	// pair_<chainId>:<pairAddress>
	ID string `json:"id"`
	// token0 is the base token
	IsBaseToken bool       `json:"isBaseToken"`
	QuoteToken  QuoteToken `json:"quoteToken"`
	V2          *V2        `json:"v2,omitempty"`
}

type Base struct {
	// chain id
	ChainID  int64  `json:"chainId"`
	Decimals *int64 `json:"decimals,omitempty"`
	// dex id
	DexID   string  `json:"dexId"`
	Factory *string `json:"factory,omitempty"`
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
}

type BaseToken struct {
	Address string `json:"address"`
	PythID  string `json:"pythId"`
	TokenID string `json:"tokenId"`
}

type Block struct {
	// transaction from address
	From             string `json:"from"`
	Hash             string `json:"hash"`
	LogIndex         int64  `json:"logIndex"`
	Number           int64  `json:"number"`
	Timestamp        int64  `json:"timestamp"`
	TransactionIndex int64  `json:"transactionIndex"`
}

type QuoteToken struct {
	Address string `json:"address"`
	PythID  string `json:"pythId"`
	TokenID string `json:"tokenId"`
}

type V2 struct {
	Swap        *Swap   `json:"swap,omitempty"`
	Sync        *Sync   `json:"sync,omitempty"`
	TotalSupply *string `json:"totalSupply,omitempty"`
}

type Swap struct {
	Amount0In  string `json:"amount0In"`
	Amount0Out string `json:"amount0Out"`
	Amount1In  string `json:"amount1In"`
	Amount1Out string `json:"amount1Out"`
	To         string `json:"to"`
}

type Sync struct {
	Reserve0 string `json:"reserve0"`
	Reserve1 string `json:"reserve1"`
}

// end pair schema
