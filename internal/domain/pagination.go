package domain

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func NewPagination(limit, page int) *Pagination {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	return &Pagination{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}
