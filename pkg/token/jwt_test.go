package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	// Sample test keys (replace with your actual keys)
	testPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

	testPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0RM0yhpLI3ZIU4X1ezkg
RAUzPMlGY6Jwv6t84m7F4h7c4fVzJathzOWO9SumND0odBgAoNSCG9Z9MvkyHvDd
vv6p8AoqnfagXEXlfNOdtrdr6HYoqTzvcDgpi3zfzBLs0izorSzPCB+yrz1liAIX
XsUu86ERyWqYVkk2fcvj6dAJsB3yLB9wlAj862TilL00YdvgSA8khjwkIWn+acT+
5kuJBVN7+n+AJipLyGG62BkI/WZdAlxda4CHmyYJy5UDfsjTe49OTiE+RY2FMPXN
KrZEim3syba8aDsPv8oEU/4K4w0a/60VcnyzdulfRhkT9C9aoWymrtXUIgHoNwNj
eNm+3zvX0vDMj+feOfGYc3mK5+nDiiTouEu/601GIShoo7XVfi7lXyBFSx9uiKs8
097atQisCgmzh8n3a7pICRezO4/3PDGdWcdbR4PwQRy9AWOvqiMoALQ7EbyiL/ug
TljofuJ2G3rK6Q1PDmjsZIa1nWGjCPGSZ3ze01GdlMCXut+QUWIn/kdJ5/PLZaon
n0IR73iObZ5Hnwg+M40OcygA1Oa/FAlNwrQ3FY+jd3vZXEFRebIUgh8hgIEUx38X
7WlmicGeW+i5zZhCFNE7RAwsMrM6K8ZnQhVt37jOOJTx7s53xAC/7PPG9DM+nff8
ztstRl/5LyCKwq988eBcty0CAwEAAQ==
-----END PUBLIC KEY-----
`)
)

func TestNewJwt(t *testing.T) {
	// Test successful creation
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if jwtInstance == nil || err != nil {
		t.Fatal("NewJwt returned nil")
	}
	_, _err := NewJwt(JwtConfig{}) // Empty config should cause a panic
	if _err == nil {
		t.Error("NewJwt should return an error with an empty config")
	}
}

func TestSign(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	userClaim := &UserClaim{
		User: User{
			Id:    uuid.New(),
			Name:  "test user",
			Email: "test@example.com",
			Roles: []string{"user"},
		},
		ExpiresIn: time.Now().Add(time.Hour).Unix(),
	}

	tokenString, err := jwtInstance.Sign(userClaim)
	if err != nil {
		t.Fatal(err)
	}

	// Basic token structure check
	if _, err := jwtInstance.Parse(tokenString); err != nil {
		t.Errorf("Invalid token format: %v", err)
	}
}

func TestParse(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	// Valid token
	userClaim := &UserClaim{
		User:      User{Id: uuid.New()},
		ExpiresIn: time.Now().Add(time.Hour).Unix(),
	}
	tokenString, _ := jwtInstance.Sign(userClaim)
	_, err = jwtInstance.Parse(tokenString)
	if err != nil {
		t.Errorf("Unexpected error parsing valid token: %v", err)
	}

	// Invalid token (triggering the error branch in Parse)
	invalidToken := "this.is.not.a.valid.jwt"
	_, err = jwtInstance.Parse(invalidToken)
	if err == nil {
		t.Error("Expected an error parsing invalid token, but got nil")
	}
}
func TestVerify(t *testing.T) {
	cases := []struct {
		name      string
		claim     *UserClaim
		expectErr bool
		isValid   bool
	}{
		{
			name: "Default",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
			isValid:   true,
		},
		{
			name: "Invalid Id",
			claim: &UserClaim{
				User:      User{Id: uuid.Nil},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
			isValid:   false,
		},
		{
			name: "Expired",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(-time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
				},
			},
			expectErr: true,
			isValid:   false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
			if err != nil {
				t.Fatal(err)
			}
			tokenString, _ := jwtInstance.Sign(tc.claim)
			valid, err := jwtInstance.Verify(tokenString)
			if (err != nil) != tc.expectErr {
				t.Errorf("Verify() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if valid != tc.isValid {
				t.Errorf("Verify() = %v, want %v", valid, tc.isValid)
			}
		})
	}
}

func TestVerifyAndParse(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()

	testCases := []struct {
		name         string
		claim        *UserClaim
		expectErr    bool
		wantClaimNil bool
	}{
		{
			name: "Valid Token",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: now.Add(time.Hour).Unix(),
			},
			expectErr:    false,
			wantClaimNil: false,
		},
		{
			name: "Expired Token",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: now.Add(-time.Hour).Unix(),
			},
			expectErr:    true,
			wantClaimNil: false,
		},
		{
			name: "Invalid Claims",
			claim: &UserClaim{
				User:      User{Id: uuid.Nil},
				ExpiresIn: now.Add(-time.Hour).Unix(),
			},
			expectErr:    true,
			wantClaimNil: false,
		},
		{
			name: "Invalid Claims And Valid Token",
			claim: &UserClaim{
				User:      User{Id: uuid.Nil},
				ExpiresIn: now.Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
				},
			},
			expectErr:    false,
			wantClaimNil: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenString := ""
			if tc.claim != nil {
				var err error
				tokenString, err = jwtInstance.Sign(tc.claim)
				if err != nil {
					t.Fatal(err)
				}
			}

			claim, err := jwtInstance.VerifyAndParse(tokenString)
			if (err != nil) != tc.expectErr {
				t.Errorf("VerifyAndParse() error = %v, expectErr = %v", err, tc.expectErr)
			}

			if !tc.expectErr && (claim == nil && !tc.wantClaimNil) {
				t.Error("VerifyAndParse() should return a claim for a valid token, but got nil")
			}
		})
	}
}

func TestGetClaims(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}

	userClaim := &UserClaim{
		User:      User{Id: uuid.New()},
		ExpiresIn: time.Now().Add(time.Hour).Unix(),
	}
	tokenString, _ := jwtInstance.Sign(userClaim)

	// Valid token
	claims, err := jwtInstance.GetClaims(tokenString)
	if err != nil {
		t.Errorf("Unexpected error getting claims from valid token: %v", err)
	}
	if claims == nil {
		t.Error("GetClaims() should return claims for a valid token, but got nil")
	}

	// Invalid token
	invalidToken := "this.is.not.a.valid.jwt"
	claims, err = jwtInstance.GetClaims(invalidToken)
	if err == nil {
		t.Error("Expected an error getting claims from invalid token, but got nil")
	}
	if claims != nil {
		t.Error("GetClaims() should return nil claims for an invalid token")
	}

	// Additional test case to cover the "return nil, nil" branch
	jwtInstance, err = NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	token, _ := jwtInstance.SignWithJWtClaims(jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	claims, err = jwtInstance.GetClaims(token)
	if err != nil {
		t.Errorf("Unexpected error getting claims from valid token: %v", err)
	}
	if claims != nil {
		t.Error("GetClaims() should return nil claims for an invalid token")
	}
}

func TestRefresh(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name      string
		claim     *UserClaim
		expectErr bool
	}{
		{
			name: "Default",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
		},
		{
			name: "Invalid Id",
			claim: &UserClaim{
				User:      User{Id: uuid.Nil},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
		},
		{
			name: "Expired",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(-time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tokenString, _ := jwtInstance.Sign(tc.claim)
			_, err := jwtInstance.Refresh(tokenString, time.Hour, nil)
			if (err != nil) != tc.expectErr {
				t.Errorf("Verify() error = %v, expectErr = %v", err, tc.expectErr)
			}

		})
	}
}

func TestExpire(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name      string
		claim     *UserClaim
		expectErr bool
	}{
		{
			name: "Default",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
		},
		{
			name: "Invalid Id",
			claim: &UserClaim{
				User:      User{Id: uuid.Nil},
				ExpiresIn: time.Now().Add(time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			},
			expectErr: false,
		},
		{
			name: "Expired",
			claim: &UserClaim{
				User:      User{Id: uuid.New()},
				ExpiresIn: time.Now().Add(-time.Hour).Unix(),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tokenString, _ := jwtInstance.Sign(tc.claim)
			_, err := jwtInstance.Expire(tokenString)
			if (err != nil) != tc.expectErr {
				t.Errorf("Verify() error = %v, expectErr = %v", err, tc.expectErr)
			}

		})
	}
}

// ... similar tests for VerifyAndParse, GetClaims, Refresh, Expire, customLogic

func TestParseWithInvalidKey(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}
	userClaim := &UserClaim{
		User:      User{Id: uuid.New()},
		ExpiresIn: time.Now().Add(time.Hour).Unix(), // Valid token
	}
	_, _ = jwtInstance.Sign(userClaim)

	// Modify the public key (making it invalid)
	invalidJwtInstance, _err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: []byte("invalid key")})

	if _err == nil {
		t.Error("Expected an error when creating Jwt with invalid key, but got nil")
	}

	if invalidJwtInstance != nil {
		t.Error("Jwt instance should be nil when created with invalid key")
	}
}

func TestCustomLogic(t *testing.T) {
	jwtInstance, err := NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey})
	if err != nil {
		t.Fatal(err)
	}

	// Test with valid signing method
	token := jwt.New(jwt.SigningMethodRS256)
	if err := jwtInstance.customLogic(token); err != nil {
		t.Errorf("Unexpected error with valid signing method: %v", err)
	}

	// Test with invalid signing method
	token = jwt.New(jwt.SigningMethodHS256)
	if err := jwtInstance.customLogic(token); err == nil {
		t.Error("Expected an error with invalid signing method, but got nil")
	}

	// validate custom logic fail
	jwtInstance, err = NewJwt(JwtConfig{PrivateKey: testPrivateKey, PublicKey: testPublicKey, SignMethod: "RS512"})
	if err != nil {
		t.Fatal(err)
	}
	tokenString, _ := jwtInstance.SignWithJWtClaims(jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	valid, err := jwtInstance.Verify(tokenString)
	if err == nil { // Expect an error for an expired token
		t.Error("Expired token should return an error, but didn't")
	}
	if valid {
		t.Error("Expired token incorrectly reported as valid")
	}
}
