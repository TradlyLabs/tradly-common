package vredis

// A geographical coordinate
type Pair struct {
	// Pair address
	Address string `json:"address"`
	// <appId>:<pairAddress>
	AppPairID string    `json:"appPairId"`
	Base      *Base     `json:"base,omitempty"`
	BaseToken BaseToken `json:"baseToken"`
	Block     *Block    `json:"block,omitempty"`
	// token0 is the base token
	IsBaseToken bool   `json:"isBaseToken"`
	Order       *Order `json:"order,omitempty"`
	// pair_<chainId>:<pairAddress>
	PairID     string     `json:"pairId"`
	QuoteToken QuoteToken `json:"quoteToken"`
	V2         *V2        `json:"v2,omitempty"`
}

type Base struct {
	// chain id
	ChainID  int64   `json:"chainId"`
	Decimals *int64  `json:"decimals,omitempty"`
	Factory  *string `json:"factory,omitempty"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
}

type BaseToken struct {
	Address     string `json:"address"`
	Decimals    int64  `json:"decimals"`
	MaxSupply   int64  `json:"maxSupply"`
	PairTokenID string `json:"pairTokenId"`
	PythID      string `json:"pythId"`
	TotalSupply int64  `json:"totalSupply"`
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

type Order struct {
	Account   string `json:"account"`
	AppID     string `json:"appId"`
	DexID     string `json:"dexId"`
	OrderID   string `json:"orderId"`
	Timestamp int64  `json:"timestamp"`
}

type QuoteToken struct {
	Address     string `json:"address"`
	Decimals    int64  `json:"decimals"`
	MaxSupply   int64  `json:"maxSupply"`
	PairTokenID string `json:"pairTokenId"`
	PythID      string `json:"pythId"`
	TotalSupply int64  `json:"totalSupply"`
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
