package aggregates

import (
	"testing"
	"time"

	"github.com/9ssi7/gopre/internal/domain/valobj"
	"github.com/stretchr/testify/assert"
)

func TestSession_SetFromDevice(t *testing.T) {
	session := &Session{}
	device := &valobj.Device{
		Name: "Test Device",
		Type: "mobile",
		OS:   "Android",
		IP:   "127.0.0.1",
	}

	session.SetFromDevice(device)

	assert.Equal(t, device.Name, session.DeviceName)
	assert.Equal(t, device.Type, session.DeviceType)
	assert.Equal(t, device.OS, session.DeviceOS)
	assert.Equal(t, device.IP, session.IpAddress)
}

func TestSession_IsRefreshValid(t *testing.T) {
	session := &Session{
		RefreshToken: "refresh_token",
		AccessToken:  "access_token",
		IpAddress:    "127.0.0.1",
	}

	testCases := []struct {
		name         string
		accessToken  string
		refreshToken string
		ipAddress    string
		expected     bool
	}{
		{"Valid", "access_token", "refresh_token", "127.0.0.1", true},
		{"Invalid Access Token", "wrong_access_token", "refresh_token", "127.0.0.1", false},
		{"Invalid Refresh Token", "access_token", "wrong_refresh_token", "127.0.0.1", false},
		{"Invalid IP Address", "access_token", "refresh_token", "192.168.0.1", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := session.IsRefreshValid(tc.accessToken, tc.refreshToken, tc.ipAddress)
			assert.Equal(t, tc.expected, isValid)
		})
	}
}

func TestSession_IsAccessValid(t *testing.T) {
	session := &Session{
		AccessToken: "access_token",
		IpAddress:   "127.0.0.1",
	}

	testCases := []struct {
		name        string
		accessToken string
		ipAddress   string
		expected    bool
	}{
		{"Valid", "access_token", "127.0.0.1", true},
		{"Invalid Access Token", "wrong_access_token", "127.0.0.1", false},
		{"Invalid IP Address", "access_token", "192.168.0.1", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := session.IsAccessValid(tc.accessToken, tc.ipAddress)
			assert.Equal(t, tc.expected, isValid)
		})
	}
}
func TestSession_Refresh(t *testing.T) {
	session := &Session{
		AccessToken: "old_access_token",
		LastLogin:   time.Now().Add(-1 * time.Hour),
		UpdatedAt:   time.Now().Add(-1 * time.Hour),
	}

	newToken := "new_access_token"
	session.Refresh(newToken)

	assert.Equal(t, newToken, session.AccessToken)
	assert.WithinDuration(t, time.Now(), session.LastLogin, 1*time.Second)
	assert.WithinDuration(t, time.Now(), session.UpdatedAt, 1*time.Second)
}

func TestSession_VerifyToken(t *testing.T) {
	session := &Session{AccessToken: "access_token"}

	assert.True(t, session.VerifyToken("access_token"))
	assert.False(t, session.VerifyToken("wrong_token"))
}

func TestSession_VerifyRefreshToken(t *testing.T) {
	session := &Session{RefreshToken: "refresh_token"}

	assert.True(t, session.VerifyRefreshToken("refresh_token"))
	assert.False(t, session.VerifyRefreshToken("wrong_token"))
}

func TestSession_SetFcmToken(t *testing.T) {
	session := &Session{}
	fcmToken := "fcm_token_123"

	session.SetFcmToken(fcmToken)
	assert.Equal(t, fcmToken, session.FcmToken)
}

func TestNewSession(t *testing.T) {
	device := valobj.Device{
		Name: "Test Device",
		Type: "mobile",
		OS:   "Android",
		IP:   "127.0.0.1",
	}
	deviceId := "device123"
	accessToken := "access_token"
	refreshToken := "refresh_token"

	session := NewSession(device, deviceId, accessToken, refreshToken)

	assert.Equal(t, deviceId, session.DeviceId)
	assert.Equal(t, device.Name, session.DeviceName)
	assert.Equal(t, device.Type, session.DeviceType)
	assert.Equal(t, device.OS, session.DeviceOS)
	assert.Equal(t, device.IP, session.IpAddress)
	assert.Equal(t, accessToken, session.AccessToken)
	assert.Equal(t, refreshToken, session.RefreshToken)
	assert.WithinDuration(t, time.Now(), session.LastLogin, 1*time.Second)
	assert.WithinDuration(t, time.Now(), session.CreatedAt, 1*time.Second)
	assert.WithinDuration(t, time.Now(), session.UpdatedAt, 1*time.Second)
}
