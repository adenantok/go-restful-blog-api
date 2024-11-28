package utils

import "fmt"

// PaginationMeta adalah struktur metadata untuk pagination
type PaginationMeta struct {
	CurrentPage  int    `json:"current_page"`
	ItemsPerPage int    `json:"items_per_page"`
	TotalItems   int    `json:"total_items"`
	TotalPages   int    `json:"total_pages"`
	HasNext      bool   `json:"has_next"`
	HasPrev      bool   `json:"has_prev"`
	NextPage     string `json:"next_page"`
	PrevPage     string `json:"prev_page"`
}

// GeneratePaginationMeta membuat metadata pagination dari total records, page dan limit
func GeneratePaginationMeta(totalRecords, page, limit int) PaginationMeta {
	totalPages := totalRecords / limit
	if totalRecords%limit != 0 {
		totalPages++
	}

	hasNext := page < totalPages
	hasPrev := page > 1

	nextPage := ""
	if hasNext {
		nextPage = fmt.Sprintf("/posts?page=%d", page+1)
	}

	prevPage := ""
	if hasPrev {
		prevPage = fmt.Sprintf("/posts?page=%d", page-1)
	}

	return PaginationMeta{
		CurrentPage:  page,
		ItemsPerPage: limit,
		TotalItems:   totalRecords,
		TotalPages:   totalPages,
		HasNext:      hasNext,
		HasPrev:      hasPrev,
		NextPage:     nextPage,
		PrevPage:     prevPage,
	}
}
