package aggregates

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVerify_IsExpired(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name      string
		expiresAt int64
		expected  bool
	}{
		{"Expired", now.Add(-1 * time.Minute).Unix(), true},
		{"Not Expired", now.Add(1 * time.Minute).Unix(), false},
		{"Expires Now", now.Unix(), false}, // Edge case: expires exactly at the current time
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := &Verify{ExpiresAt: tc.expiresAt}
			assert.Equal(t, tc.expected, v.IsExpired())
		})
	}
}

func TestVerify_IsExceeded(t *testing.T) {
	testCases := []struct {
		name     string
		tryCount int
		expected bool
	}{
		{"Exceeded", 3, true},
		{"Not Exceeded", 2, false},
		{"At Limit", 2, false}, // Edge case: at the limit, but not exceeded
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := &Verify{TryCount: tc.tryCount}
			assert.Equal(t, tc.expected, v.IsExceeded())
		})
	}
}

func TestVerify_IncTryCount(t *testing.T) {
	v := &Verify{TryCount: 1}
	v.IncTryCount()
	assert.Equal(t, 2, v.TryCount)
}

func TestNewVerify(t *testing.T) {
	userId := uuid.New()
	deviceId := "device123"
	locale := "en-US"
	v := NewVerify(userId, deviceId, locale)

	assert.Equal(t, userId, v.UserId)
	assert.Equal(t, deviceId, v.DeviceId)
	assert.Equal(t, locale, v.Locale)
	assert.Len(t, v.Code, 4) // Code should be 4 digits
	assert.Equal(t, 0, v.TryCount)
	assert.WithinDuration(t, time.Now().Add(5*time.Minute), time.Unix(v.ExpiresAt, 0), 1*time.Second)
}
