package token

import (
	"os"
	"time"
)

type service struct {
	jwt *Jwt
	cnf Config
}

type Config struct {
	PublicKeyFile  string
	PrivateKeyFile string
	Project        string
	SignMethod     string
}

const (
	AccessTokenDuration  time.Duration = time.Hour * 24
	RefreshTokenDuration time.Duration = time.Hour * 24 * 30
)

func New(cnf Config) (*service, error) {
	j, err := NewJwt(JwtConfig{
		PublicKey:  readFile(cnf.PublicKeyFile),
		PrivateKey: readFile(cnf.PrivateKeyFile),
		SignMethod: cnf.SignMethod,
	})
	if err != nil {
		return nil, err
	}
	return &service{
		jwt: j,
		cnf: cnf,
	}, nil
}

func readFile(name string) []byte {
	f, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return f
}

func (t *service) GenerateAccessToken(u User) (string, error) {
	claims := &UserClaim{
		User:      u,
		Project:   t.cnf.Project,
		IsAccess:  true,
		IsRefresh: false,
	}
	claims.SetExpireIn(AccessTokenDuration)
	return t.generate(claims)
}

func (t *service) GenerateRefreshToken(u User) (string, error) {
	claims := &UserClaim{
		User:      u,
		Project:   t.cnf.Project,
		IsAccess:  false,
		IsRefresh: true,
	}
	claims.SetExpireIn(RefreshTokenDuration)
	return t.generate(claims)
}

func (t *service) generate(u *UserClaim) (string, error) {
	tkn, err := t.jwt.Sign(u)
	if err != nil {
		return "", err
	}
	return tkn, nil
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
