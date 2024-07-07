package list

// PagiRequest is a struct for pagination request
type PagiRequest struct {

	// Page is a pointer to int
	// Page is a query parameter for page number
	// Page is validated to be greater than 0
	Page *int `query:"page" validate:"omitempty,gt=0"`

	// Limit is a pointer to int
	// Limit is a query parameter for limit number of items per page
	// Limit is validated to be greater than 0
	Limit *int `query:"limit" validate:"omitempty,gt=0"`
}

// Default is a method to set default value for PagiRequest
// Default set Page to 1 if Page is nil or less than 1
// Default set Limit to 10 if Limit is nil or less than 1
func (r *PagiRequest) Default() {
	if r.Page == nil || *r.Page <= 0 {
		r.Page = new(int)
		*r.Page = 1
	}
	if r.Limit == nil || *r.Limit <= 0 {
		r.Limit = new(int)
		*r.Limit = 10
	}
}

// Offset is a method to calculate offset for pagination
func (r *PagiRequest) Offset() int {
	return (*r.Page - 1) * *r.Limit
}

// TotalPage is a method to calculate total page for pagination
func (r *PagiRequest) TotalPage(total int64) int {
	if total == 0 {
		return 0
	}
	return int(total)/(*r.Limit) + 1
}
