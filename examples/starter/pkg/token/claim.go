package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone *string   `json:"phone"`
	Roles []string  `json:"roles"`
}

type UserClaim struct {
	User
	jwt.RegisteredClaims
	ExpiresIn int64  `json:"expiresIn"`
	Project   string `json:"project"`
	IsAccess  bool   `json:"isAccess"`
	IsRefresh bool   `json:"isRefresh"`
}

func (c *UserClaim) Valid() error {
	if c.IsExpired() {
		return errors.New("user_expired")
	}
	return nil
}

func (c *UserClaim) Expire() {
	c.ExpiresIn = time.Now().Unix() - 1
}

func (c *UserClaim) SetExpireIn(d time.Duration) {
	c.ExpiresIn = time.Now().Add(d).Unix()
}

func (c *UserClaim) IsExpired() bool {
	return c.ExpiresIn < time.Now().Unix()
}
