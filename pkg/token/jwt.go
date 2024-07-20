package token

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Jwt struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	signMethod string
}

type JwtConfig struct {
	PublicKey  []byte
	PrivateKey []byte
	SignMethod string
}

func NewJwt(config JwtConfig) (*Jwt, error) {
	if config.SignMethod == "" {
		config.SignMethod = "RS256"
	}
	priv, pub, err := parseKeys(config)
	if err != nil {
		return nil, err
	}
	return &Jwt{
		publicKey:  pub,
		privateKey: priv,
		signMethod: config.SignMethod,
	}, nil
}

func parseKeys(config JwtConfig) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := parsePrivateKey(config.PrivateKey)
	if err != nil {
		return nil, nil, err
	}
	public, err := parsePublicKey(config.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return private, public, nil
}

func parsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func parsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (j *Jwt) Sign(p *UserClaim) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(j.signMethod))
	token.Claims = p
	return token.SignedString(j.privateKey)
}

func (j *Jwt) SignWithJWtClaims(p jwt.Claims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(j.signMethod))
	token.Claims = p
	return token.SignedString(j.privateKey)
}

func (j *Jwt) Parse(t string, options ...jwt.ParserOption) (*jwt.Token, error) {
	return jwt.ParseWithClaims(t, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if err := j.customLogic(token); err != nil {
			return nil, err
		}
		return j.publicKey, nil
	}, options...)
}

func (j *Jwt) Verify(t string) (bool, error) {
	token, err := j.Parse(t)
	if err != nil {
		return false, err
	}
	if res, ok := token.Claims.(*UserClaim); ok && res.Id != uuid.Nil {
		return true, nil
	}
	return false, nil
}

func (j *Jwt) VerifyAndParse(t string) (*UserClaim, error) {
	token, err := j.Parse(t)
	if err != nil {
		return nil, err
	}
	if res, ok := token.Claims.(*UserClaim); ok && res.Id != uuid.Nil {
		return res, nil
	}
	return nil, nil
}

func (j *Jwt) customLogic(token *jwt.Token) error {
	if token.Method != jwt.SigningMethodRS256 {
		return fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return nil
}

func (j *Jwt) GetClaims(t string) (*UserClaim, error) {
	token, err := j.Parse(t, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}
	if res, ok := token.Claims.(*UserClaim); ok && res.Id != uuid.Nil {
		return res, nil
	}
	return nil, nil
}

func (j *Jwt) Refresh(t string, d time.Duration, nd *jwt.NumericDate) (string, error) {
	token, err := j.Parse(t)
	if err != nil {
		return "", err
	}
	if res, ok := token.Claims.(*UserClaim); ok && res.Id != uuid.Nil {
		res.ExpiresIn = time.Now().Add(d).Unix()
		return j.Sign(res)
	}
	return "", nil
}

func (j *Jwt) Expire(t string) (string, error) {
	token, err := j.Parse(t)
	if err != nil {
		return "", err
	}
	if res, ok := token.Claims.(*UserClaim); ok && res.Id != uuid.Nil {
		res.ExpiresIn = time.Now().Add(-time.Hour).Unix()
		res.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(-time.Hour))
		return j.Sign(res)
	}
	return "", nil
}
