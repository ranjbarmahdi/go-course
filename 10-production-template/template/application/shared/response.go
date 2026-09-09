package shared

type PageMeta struct {
	Page       int
	PerPage    int
	TotalCount int
	PageCount  int
}

// NewPageMeta derives response metadata from the Page in request.go and the
// total row count reported by the repository.
func NewPageMeta(p Page, totalCount int) PageMeta {
	pageCount := 0
	if p.Size > 0 {
		pageCount = (totalCount + p.Size - 1) / p.Size
	}

	return PageMeta{
		Page:       p.Number,
		PerPage:    p.Size,
		TotalCount: totalCount,
		PageCount:  pageCount,
	}
}

type PageResult[T any] struct {
	Items []T
	Meta  PageMeta
}

func NewPageResult[T any](items []T, p Page, totalCount int) PageResult[T] {
	if items == nil {
		items = []T{}
	}

	return PageResult[T]{Items: items, Meta: NewPageMeta(p, totalCount)}
}
