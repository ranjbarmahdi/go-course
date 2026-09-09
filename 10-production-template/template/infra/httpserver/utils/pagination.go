package utils

import (
	"net/http"
	"strconv"
	"template/application/shared"
)

const (
	DefaultPage    = 1
	DefaultPerPage = 50
	MaxPerPage     = 100
)

// ParsePage mirrors the gateway's pagination pipe: missing, malformed, or
// oversized values fall back to defaults so a list endpoint always answers.
func ParsePage(r *http.Request, defaultPerPage, maxPerPage int) shared.Page {
	q := r.URL.Query()
	page := shared.Page{Number: DefaultPage, Size: defaultPerPage}

	if n, err := strconv.Atoi(q.Get("page")); err == nil && n > 0 {
		page.Number = n
	}

	if s, err := strconv.Atoi(q.Get("per_page")); err == nil && s > 0 {
		page.Size = min(s, maxPerPage)
	}

	return page
}

type PageMetaResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
	PageCount  int `json:"page_count"`
}

type PageResponse[T any] struct {
	Items []T              `json:"items"`
	Meta  PageMetaResponse `json:"meta"`
}

// ToApiPageResponse converts an application page result to its wire form. It is
// a function rather than a method because Go methods cannot be generic.
func ToApiPageResponse[A, R any](res shared.PageResult[A], toItem func(A) R) PageResponse[R] {
	items := make([]R, 0, len(res.Items))
	for _, item := range res.Items {
		items = append(items, toItem(item))
	}

	return PageResponse[R]{
		Items: items,
		Meta: PageMetaResponse{
			Page:       res.Meta.Page,
			PerPage:    res.Meta.PerPage,
			TotalCount: res.Meta.TotalCount,
			PageCount:  res.Meta.PageCount,
		},
	}
}
