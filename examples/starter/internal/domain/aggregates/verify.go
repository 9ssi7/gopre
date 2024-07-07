package aggregates

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Verify struct {
	DeviceId  string    `json:"device_id"`
	Locale    string    `json:"locale"`
	UserId    uuid.UUID `json:"user_id"`
	Code      string    `json:"code"`
	TryCount  int       `json:"try_count"`
	ExpiresAt int64     `json:"expires_at"`
}

func (v *Verify) IsExpired() bool {
	return v.ExpiresAt < time.Now().Unix()
}

func (v *Verify) IsExceeded() bool {
	return v.TryCount >= 3
}

func (v *Verify) IncTryCount() {
	v.TryCount++
}

func NewVerify(userId uuid.UUID, deviceId string, locale string) *Verify {
	return &Verify{
		UserId:    userId,
		DeviceId:  deviceId,
		Locale:    locale,
		Code:      fmt.Sprintf("%04d", rand.Intn(9999)),
		TryCount:  0,
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	}
}
