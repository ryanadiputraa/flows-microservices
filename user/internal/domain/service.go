package domain

type ServiveResponse[T any] struct {
	Message string              `json:"message"`
	ErrCode string              `json:"err_code"`
	Errors  map[string][]string `json:"errros"`
	Data    T                   `json:"data"`
}
