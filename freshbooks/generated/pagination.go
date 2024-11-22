package generated

import "time"

type SortBy struct {
	Asc  bool
	Name string

	AfterDate    *time.Time
	UpdatedAfter bool
}

var defaultSortBy = &SortBy{
	Asc:  false,
	Name: "date",
}

func DefaultSortBy() *SortBy {
	return defaultSortBy
}

type Pagination struct {
	Page    uint64 `json:"page"`
	Pages   uint64 `json:"pages"`
	PerPage uint64 `json:"per_page"`
	Total   uint64 `json:"total"`
}

func ListAll[T any](id string, sortBy *SortBy, lister func(id string, page uint64, sortBy *SortBy) ([]T, *Pagination, error)) ([]T, error) {
	results := make([]T, 0, 10)

	result, pagination, err := lister(id, 1, sortBy)
	results = append(results, result...)
	if err != nil {
		return results, err
	}

	for ; err == nil && len(result) > 0 && pagination.Page != pagination.Pages; result, pagination, err = lister(id, pagination.Page+1, sortBy) {
		results = append(results, result...)
	}

	if err != nil {
		return results, err
	}

	return results, err
}
