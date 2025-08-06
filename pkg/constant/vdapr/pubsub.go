package vdapr

import "github.com/dapr/go-sdk/service/common"

var (
	SUB_UNISWAP_V2_PAIRS = &common.Subscription{
		PubsubName: "management",
		Topic:      "uniswap-v2-pairs",
		Route:      "/uniswap-v2-pairs",
	}
	SUB_PRICE_FEEDS = &common.Subscription{
		PubsubName: "management",
		Topic:      "price-feeds",
		Route:      "/price-feeds",
	}
	SUB_BOT = &common.Subscription{
		PubsubName: "management",
		Topic:      "bot",
		Route:      "/bot",
	}
)
