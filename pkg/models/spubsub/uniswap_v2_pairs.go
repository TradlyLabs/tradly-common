package spubsub

import "github.com/TemoreIO/temore-common/pkg/models/chain_data"

type UniswapV2PairEventType int

const (
	UNISWAP_V2_PAIR_EVENT_SYNC UniswapV2PairEventType = iota
	UNISWAP_V2_PAIR_EVENT_ADD
	UNISWAP_V2_PAIR_EVENT_REMOVE
)

type UniswapV2PairEvent struct {
	Type  UniswapV2PairEventType `json:"type"`
	Pairs []*chain_data.Pair     `json:"pairs"`
}
