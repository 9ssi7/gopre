package aggregates

import (
	"time"

	"github.com/9ssi7/gopre-starter/internal/domain/valobj"
)

type Session struct {
	DeviceId     string    `json:"device_id"`
	DeviceName   string    `json:"device_name"`
	DeviceType   string    `json:"device_type"`
	DeviceOS     string    `json:"device_os"`
	IpAddress    string    `json:"ip_address"`
	FcmToken     string    `json:"fcm_token"`
	RefreshToken string    `json:"refresh_token"`
	AccessToken  string    `json:"access_token"`
	LastLogin    time.Time `json:"last_login"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (s *Session) SetFromDevice(d *valobj.Device) {
	s.DeviceName = d.Name
	s.DeviceType = d.Type
	s.DeviceOS = d.OS
	s.IpAddress = d.IP
}

func (s *Session) IsRefreshValid(accessToken, refreshToken string, ipAddress string) bool {
	return s.RefreshToken == refreshToken && s.AccessToken == accessToken && s.IpAddress == ipAddress
}

func (s *Session) IsAccessValid(accessToken, ipAddress string) bool {
	return s.AccessToken == accessToken && s.IpAddress == ipAddress
}

func (s *Session) Refresh(token string) {
	s.AccessToken = token
	s.LastLogin = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Session) VerifyToken(token string) bool {
	return s.AccessToken == token
}

func (s *Session) VerifyRefreshToken(token string) bool {
	return s.RefreshToken == token
}

func (s *Session) SetFcmToken(token string) {
	s.FcmToken = token
}

func NewSession(device valobj.Device, deviceId, accessToken, refreshToken string) *Session {
	t := time.Now()
	return &Session{
		DeviceId:     deviceId,
		DeviceName:   device.Name,
		DeviceType:   device.Type,
		DeviceOS:     device.OS,
		IpAddress:    device.IP,
		LastLogin:    t,
		CreatedAt:    t,
		UpdatedAt:    t,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}
