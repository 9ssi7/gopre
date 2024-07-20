package currency

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"USD", true},
		{"EUR", true},
		{"JPY", true},
		{"GBP", true},
		{"XYZ", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			isValid := IsValid(tt.code)
			if isValid != tt.valid {
				t.Errorf("IsValid(%q) = %v, want %v", tt.code, isValid, tt.valid)
			}
		})
	}
}
