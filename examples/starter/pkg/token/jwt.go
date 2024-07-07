package token

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

type JwtConfig struct {
	PublicKey  []byte
	PrivateKey []byte
}

func NewJwt(config JwtConfig) *Jwt {
	priv, pub := parseKeys(config)
	return &Jwt{
		publicKey:  pub,
		privateKey: priv,
	}
}

func parseKeys(config JwtConfig) (*rsa.PrivateKey, *rsa.PublicKey) {
	private, err := parsePrivateKey(config.PrivateKey)
	if err != nil {
		panic(err)
	}
	public, err := parsePublicKey(config.PublicKey)
	if err != nil {
		panic(err)
	}
	return private, public
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
	token := jwt.New(jwt.GetSigningMethod("RS256"))
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
	if _, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return true, nil
	}
	return false, nil
}

func (j *Jwt) VerifyAndParse(t string) (*UserClaim, error) {
	token, err := j.Parse(t)
	if err != nil {
		return nil, err
	}
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
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
	if res, ok := token.Claims.(*UserClaim); ok {
		return res, nil
	}
	return nil, nil
}

func (j *Jwt) Refresh(t string, d time.Duration, nd *jwt.NumericDate) (string, error) {
	token, err := j.Parse(t)
	if err != nil {
		return "", err
	}
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
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
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
		res.ExpiresIn = time.Now().Add(-time.Hour).Unix()
		res.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(-time.Hour))
		return j.Sign(res)
	}
	return "", nil
}
