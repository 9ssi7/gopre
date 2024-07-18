package list

// PagiResponse is a struct for pagination response
// PagiResponse contains page number, limit, total items, filtered total items, total page, and list of items
// PagiResponse is a generic struct
type PagiResponse[T any] struct {
	Page          int   `json:"page"`
	Limit         int   `json:"limit"`
	Total         int64 `json:"total"`
	FilteredTotal int64 `json:"filtered_total"`
	TotalPage     int   `json:"total_page"`
	List          []T   `json:"list"`
}
