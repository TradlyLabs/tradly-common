package vredis

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/TradlyLabs/tradly-common/pkg/w3utils"
)

func randomAddress(length int) string {

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length-2)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return "0x" + string(result)

}

func randomId() string {
	return string(w3utils.GenerateID("0", 97, randomAddress(42)))
}

func generateMockData() Pair {
	decimals := int64(18)
	base := &Base{
		ChainID:  1,
		Decimals: &decimals,
		Name:     "Ether",
		Symbol:   "ETH",
	}
	basetoken := BaseToken{
		Address: randomAddress(42),
		PythID:  "0x2f95862b045670cd22bee3114c39763a4a08beeb663b145d283c31d7d1101c4f",
		TokenID: randomId(),
	}
	block := &Block{
		From:             randomAddress(42),
		Hash:             randomAddress(66),
		LogIndex:         rand.Int63(),
		Number:           rand.Int63(),
		Timestamp:        time.Now().Unix(),
		TransactionIndex: rand.Int63(),
	}
	order := &Order{
		AppID:   randomId(),
		DexID:   "0xaffb39b4aa05d622e4031ae028fd2a6f9abe1ea25f32094db20b59a2f463d5db",
		OrderID: randomId(),
	}
	quotetoken := QuoteToken{
		Address: randomAddress(42),
		PythID:  randomAddress(32),
		TokenID: "token2",
	}
	swap := &Swap{
		Amount0In:  "1000000000000000000", // 1 token with 18 decimals
		Amount0Out: "0",
		Amount1In:  "0",
		Amount1Out: "2000000000000000000", // 2 tokens with 18 decimals
		To:         randomAddress(42),
	}
	sync := &Sync{
		Reserve0: "1000000000000000000", // 1 token with 18 decimals
		Reserve1: "2000000000000000000", // 2 tokens with 18 decimals
	}
	v2 := &V2{
		Swap: swap,
		Sync: sync,
	}

	pair := Pair{
		Address:     randomAddress(42),
		AppPairID:   fmt.Sprintf("appId:%s", randomAddress(42)),
		Base:        base,
		BaseToken:   basetoken,
		Block:       block,
		IsBaseToken: true,
		Order:       order,
		PairID:      fmt.Sprintf("pair_%d:%s", 1, randomAddress(42)),
		QuoteToken:  quotetoken,
		V2:          v2,
	}
	return pair
}

func TestGenerateMockData(t *testing.T) {
	pair := generateMockData()
	byt, err := json.Marshal(pair)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(byt))
}
