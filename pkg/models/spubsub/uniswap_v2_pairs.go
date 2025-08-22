package spubsub

import "github.com/TradlyLabs/tradly-common/pkg/models/chain_data"

type UniswapV2PairEventType int

const (
	UNISWAP_V2_PAIR_EVENT_SYNC UniswapV2PairEventType = iota
	UNISWAP_V2_PAIR_EVENT_ADD
	UNISWAP_V2_PAIR_EVENT_REMOVE
	UNISWAP_V2_PAIR_EVENT_UPDATE
)

type UniswapV2PairEvent struct {
	Type  UniswapV2PairEventType `json:"type"`
	Pairs []*chain_data.Pair     `json:"pairs"`
}

type TokenEventType int

const (
	TOKEN_EVENT_SYNC TokenEventType = iota
	TOKEN_EVENT_ADD
	TOKEN_EVENT_REMOVE
	TOKEN_EVENT_UPDATE
	TOKEN_EVENT_MARK
)

type TokenEvent struct {
	Type   TokenEventType      `json:"type"`
	Tokens []*chain_data.Token `json:"tokens"`
}
