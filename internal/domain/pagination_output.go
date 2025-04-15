package domain

type PaginationOutput[T any] struct {
	Total       int  `json:"total"`
	CurrentPage int  `json:"current_page"`
	LastPage    int  `json:"last_page"`
	Data        []*T `json:"data"`
}
