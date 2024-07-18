package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestUserClaim_Valid(t *testing.T) {
	testCases := []struct {
		name      string
		claim     *UserClaim
		expectErr bool
	}{
		{
			name: "Valid Claim",
			claim: &UserClaim{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expires in one hour
				},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
			},
			expectErr: false,
		},
		{
			name: "Expired Claim",
			claim: &UserClaim{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // Expired an hour ago
				},
				ExpiresIn: time.Now().Add(-time.Hour).Unix(),
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.claim.Valid()
			if (err != nil) != tc.expectErr {
				t.Errorf("Valid() error = %v, expectErr = %v", err, tc.expectErr)
			}
		})
	}
}

func TestUserClaim_Expire(t *testing.T) {
	claim := &UserClaim{}
	claim.Expire()
	if !claim.IsExpired() {
		t.Error("Expire() did not set the claim to expired")
	}
}

func TestUserClaim_SetExpireIn(t *testing.T) {
	claim := &UserClaim{}
	claim.SetExpireIn(time.Hour)
	if claim.IsExpired() {
		t.Error("SetExpireIn() set the claim to expired unexpectedly")
	}
}

func TestUserClaim_IsExpired(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name      string
		expiresIn int64
		want      bool
	}{
		{
			name:      "Expired",
			expiresIn: now.Add(-time.Hour).Unix(),
			want:      true,
		},
		{
			name:      "Not Expired",
			expiresIn: now.Add(time.Hour).Unix(),
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claim := &UserClaim{ExpiresIn: tc.expiresIn}
			if got := claim.IsExpired(); got != tc.want {
				t.Errorf("IsExpired() = %v, want %v", got, tc.want)
			}
		})
	}
}
