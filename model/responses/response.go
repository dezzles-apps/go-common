package responses

type Response[T any] struct {
	Data T `json:"data"`
}
