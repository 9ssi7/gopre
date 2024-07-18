package list

import "testing"

func TestPagiRequest_Default(t *testing.T) {
	tests := []struct {
		name  string
		input PagiRequest
		want  PagiRequest
	}{
		{
			name:  "All Nil",
			input: PagiRequest{},
			want:  PagiRequest{Page: intPtr(1), Limit: intPtr(10)},
		},
		{
			name:  "Page Zero",
			input: PagiRequest{Page: intPtr(0)},
			want:  PagiRequest{Page: intPtr(1), Limit: intPtr(10)},
		},
		{
			name:  "Limit Zero",
			input: PagiRequest{Limit: intPtr(0)},
			want:  PagiRequest{Page: intPtr(1), Limit: intPtr(10)},
		},
		{
			name:  "Valid Input",
			input: PagiRequest{Page: intPtr(3), Limit: intPtr(25)},
			want:  PagiRequest{Page: intPtr(3), Limit: intPtr(25)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.Default()
			if *tt.input.Page != *tt.want.Page || *tt.input.Limit != *tt.want.Limit {
				t.Errorf("Default() got = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func TestPagiRequest_Offset(t *testing.T) {
	tests := []struct {
		name   string
		page   int
		limit  *int // limit'i pointer olarak değiştirin.
		offset int
	}{
		{name: "First Page (Nil Limit)", page: 1, limit: nil, offset: 0},
		{name: "Second Page (Nil Limit)", page: 2, limit: nil, offset: 10},
		{name: "Third Page (Nil Limit)", page: 3, limit: nil, offset: 20},
		{name: "First Page", page: 1, limit: intPtr(10), offset: 0},
		{name: "Second Page", page: 2, limit: intPtr(10), offset: 10},
		{name: "Third Page", page: 3, limit: intPtr(15), offset: 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := PagiRequest{Page: intPtr(tt.page), Limit: tt.limit}
			if got := r.Offset(); got != tt.offset {
				t.Errorf("Offset() = %v, want %v", got, tt.offset)
			}

			// Verify that Default() was called if Limit was nil
			if tt.limit == nil && *r.Limit != 10 {
				t.Errorf("Default() was not called when Limit was nil")
			}
		})
	}
}

func TestPagiRequest_TotalPage(t *testing.T) {
	tests := []struct {
		name     string
		limit    *int
		total    int64
		wantPage int
	}{
		{name: "Zero Total", limit: nil, total: 0, wantPage: 0},
		{name: "Single Page", limit: nil, total: 5, wantPage: 1},
		{name: "Multiple Pages", limit: nil, total: 25, wantPage: 3},
		{name: "Exact Multiple", limit: nil, total: 30, wantPage: 4},
		{name: "Single Page (With Limit)", limit: intPtr(10), total: 5, wantPage: 1},
		{name: "Multiple Pages (With Limit)", limit: intPtr(10), total: 25, wantPage: 3},
		{name: "Exact Multiple (With Limit)", limit: intPtr(10), total: 30, wantPage: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := PagiRequest{Limit: tt.limit}
			if got := r.TotalPage(tt.total); got != tt.wantPage {
				t.Errorf("TotalPage() = %v, want %v", got, tt.wantPage)
			}

			// Verify that Default() was called if Limit was nil
			if tt.limit == nil && tt.wantPage != 0 && *r.Limit != 10 {
				t.Errorf("Default() was not called when Limit was nil")
			}
		})
	}
}

func intPtr(v int) *int { return &v } // Helper function for creating int pointers
