package model

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

type ErrResponse struct {
	Message string `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}
