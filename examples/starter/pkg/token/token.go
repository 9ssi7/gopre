package token

import (
	"os"
	"time"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	jwt *Jwt
}

const (
	AccessTokenDuration  time.Duration = time.Hour * 24
	RefreshTokenDuration time.Duration = time.Hour * 24 * 30
)

var client *service

func Init() {
	cnf := config.ReadValue().RSA
	j := NewJwt(JwtConfig{
		PublicKey:  readFile(cnf.PublicKeyFile),
		PrivateKey: readFile(cnf.PrivateKeyFile),
	})
	client = &service{
		jwt: j,
	}
}

func Client() *service {
	return client
}

func readFile(name string) []byte {
	f, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return f
}

func (t *service) Generate(u User) (string, string, error) {
	accessToken, err := t.GenerateAccessToken(u)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := t.GenerateRefreshToken(u)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (t *service) GenerateAccessToken(u User) (string, error) {
	return t.generate(&UserClaim{
		User: u,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
		},
		ExpiresIn: time.Now().Add(AccessTokenDuration).Unix(),
		Project:   config.ReadValue().TokenSrv.Project,
		IsAccess:  true,
		IsRefresh: false,
	})
}

func (t *service) GenerateRefreshToken(u User) (string, error) {
	return t.generate(&UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
		},
		User:      u,
		ExpiresIn: time.Now().Add(RefreshTokenDuration).Unix(),
		Project:   config.ReadValue().TokenSrv.Project,
		IsAccess:  false,
		IsRefresh: true,
	})
}

func (t *service) generate(u *UserClaim) (string, error) {
	tkn, err := t.jwt.Sign(u)
	if err != nil {
		return "", err
	}
	return tkn, err
}

func (t *service) Parse(token string) (*UserClaim, error) {
	return t.jwt.GetClaims(token)
}

func (t *service) Verify(token string) (bool, error) {
	return t.jwt.Verify(token)
}

func (t *service) VerifyAndParse(token string) (*UserClaim, error) {
	return t.jwt.VerifyAndParse(token)
}
