package ginutils

type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
