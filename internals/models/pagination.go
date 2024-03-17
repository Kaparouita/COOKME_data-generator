package models

// Pagination is a struct that contains the information of the pagination
type Pagination struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	ID     int    `json:"id"`
	SortBy string `json:"sort_by"`
	Search string `json:"search"`
}
