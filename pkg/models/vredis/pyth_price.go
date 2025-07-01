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
