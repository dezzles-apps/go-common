package api

type Response[T any] struct {
	Data T `json:"data"`
}

type PageKey[T any] struct {
	Items []T     `json:"items"`
	Next  *string `json:"next"`
}

type Page[T any] struct {
	Items         []T `json:"items"`
	Page          int `json:"page"`
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
}
