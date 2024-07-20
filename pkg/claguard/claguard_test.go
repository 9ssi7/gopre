package claguard

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name       string
		ownClaims  []string
		wantClaims []string
		want       bool
	}{
		{
			name:       "No Claims",
			ownClaims:  []string{},
			wantClaims: []string{},
			want:       false,
		},
		{
			name:       "Single Claim Match",
			ownClaims:  []string{"admin"},
			wantClaims: []string{"admin"},
			want:       true,
		},
		{
			name:       "Single Claim No Match",
			ownClaims:  []string{"admin"},
			wantClaims: []string{"editor"},
			want:       false,
		},
		{
			name:       "Multiple Claims Match",
			ownClaims:  []string{"admin", "editor"},
			wantClaims: []string{"admin", "editor"},
			want:       true,
		},
		{
			name:       "Multiple Claims No Match",
			ownClaims:  []string{"admin", "editor"},
			wantClaims: []string{"viewer", "moderator"},
			want:       false,
		},
		{
			name:       "Partial Match",
			ownClaims:  []string{"admin", "editor"},
			wantClaims: []string{"admin", "viewer"},
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Check(tt.ownClaims, tt.wantClaims)
			if got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSingle(t *testing.T) {
	tests := []struct {
		name      string
		ownClaims []string
		wantClaim string
		want      bool
	}{
		{
			name:      "Empty Own Claims",
			ownClaims: []string{},
			wantClaim: "admin",
			want:      false,
		},
		{
			name:      "Single Claim Match",
			ownClaims: []string{"admin"},
			wantClaim: "admin",
			want:      true,
		},
		{
			name:      "Single Claim No Match",
			ownClaims: []string{"editor"},
			wantClaim: "admin",
			want:      false,
		},
		{
			name:      "Multiple Claims Match",
			ownClaims: []string{"admin", "editor"},
			wantClaim: "editor",
			want:      true,
		},
		{
			name:      "Multiple Claims No Match",
			ownClaims: []string{"viewer", "moderator"},
			wantClaim: "admin",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckSingle(tt.ownClaims, tt.wantClaim)
			if got != tt.want {
				t.Errorf("CheckSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}
