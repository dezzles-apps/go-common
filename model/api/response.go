package api

type Response[T any] struct {
	Data T `json:"data"`
}

type Page[T any] struct {
	Items []T     `json:"items"`
	Next  *string `json:"next"`
}
