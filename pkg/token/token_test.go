package token

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

var (
	testConfig = Config{
		PublicKeyFile:  "test_public_key.pem",
		PrivateKeyFile: "test_private_key.pem",
		Project:        "test_project",
	}
)

func TestToken(t *testing.T) {
	if err := os.WriteFile(testConfig.PublicKeyFile, testPublicKey, 0644); err != nil {
		t.Fatalf("Failed to create temp public key file: %v", err)
	}
	defer os.Remove(testConfig.PublicKeyFile)
	if err := os.WriteFile(testConfig.PrivateKeyFile, testPrivateKey, 0644); err != nil {
		t.Fatalf("Failed to create temp private key file: %v", err)
	}
	defer os.Remove(testConfig.PrivateKeyFile)
	service, err := New(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	if service == nil || service.jwt == nil {
		t.Fatal("New returned nil service or jwt")
	}

	t.Run("Invalid peys", func(t *testing.T) {
		// Test with invalid public key
		invalidConfig := testConfig
		invalidConfig.PublicKeyFile = "invalid_public_key.pem"
		if err := os.WriteFile(invalidConfig.PublicKeyFile, []byte("invalid"), 0644); err != nil {
			t.Fatalf("Failed to create temp invalid public key file: %v", err)
		}
		defer os.Remove(invalidConfig.PublicKeyFile)
		_, err = New(invalidConfig)
		if err == nil {
			t.Error("New should return an error with invalid public key, but didn't")
		}

		// Test with invalid private key
		invalidConfig = testConfig
		invalidConfig.PrivateKeyFile = "invalid_private_key.pem"
		if err := os.WriteFile(invalidConfig.PrivateKeyFile, []byte("invalid"), 0644); err != nil {
			t.Fatalf("Failed to create temp invalid private key file: %v", err)
		}
		defer os.Remove(invalidConfig.PrivateKeyFile)
		_, err = New(invalidConfig)
		if err == nil {
			t.Error("New should return an error with invalid private key, but didn't")
		}
	})

	t.Run("GenerateAccessToken", func(t *testing.T) {
		user := User{
			Id:    uuid.New(),
			Name:  "Test User",
			Email: "test@example.com",
			Roles: []string{"user"},
		}
		accessToken, err := service.GenerateAccessToken(user)
		if err != nil {
			t.Errorf("GenerateAccessToken() error = %v", err)
			return
		}
		if accessToken == "" {
			t.Fatal("GenerateAccessToken() returned an empty token")
		}
	})

	t.Run("GenerateRefreshToken", func(t *testing.T) {
		user := User{
			Id:    uuid.New(),
			Name:  "Test User",
			Email: "test@example.com",
			Roles: []string{"user"},
		}
		refreshToken, err := service.GenerateRefreshToken(user)
		if err != nil {
			t.Errorf("GenerateRefreshToken() error = %v", err)
			return
		}
		if refreshToken == "" {
			t.Fatal("GenerateRefreshToken() returned an empty token")
		}
	})

	t.Run("TokenParse", func(t *testing.T) {
		token, err := service.GenerateAccessToken(User{Id: uuid.New()})
		if err != nil {
			t.Fatal(err)
		}
		_, err = service.Parse(token)
		if err != nil {
			t.Errorf("Parse() error = %v", err)
		}
	})

	t.Run("TokenVerify", func(t *testing.T) {
		token, err := service.GenerateAccessToken(User{Id: uuid.New()})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := service.Verify(token); err != nil {
			t.Errorf("Verify() error = %v", err)
		}
	})

	t.Run("TokenVerifyAndParse", func(t *testing.T) {
		token, err := service.GenerateAccessToken(User{Id: uuid.New()})
		if err != nil {
			t.Fatal(err)
		}
		_, err = service.VerifyAndParse(token)
		if err != nil {
			t.Errorf("VerifyAndParse() error = %v", err)
		}
	})

	t.Run("generate with invalid srv", func(t *testing.T) {
		invalidSrv, err := New(Config{
			PublicKeyFile:  testConfig.PublicKeyFile,
			PrivateKeyFile: testConfig.PrivateKeyFile,
			Project:        "invalid_project",
			SignMethod:     "HS256",
		})
		if err != nil {
			t.Fatal(err)
		}
		_, err = invalidSrv.generate(&UserClaim{User: User{Id: uuid.New()}})
		if err == nil {
			t.Error("generate should return an error with invalid service, but didn't")
		}
	})

	t.Run("ReadFile", func(t *testing.T) {
		// Setup: Create a temporary file with some content
		tempFile := "test_file.txt"
		content := []byte("test content")
		if err := os.WriteFile(tempFile, content, 0644); err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tempFile)

		// Test successful read
		data := readFile(tempFile)
		if string(data) != string(content) {
			t.Errorf("readFile() = %v, want %v", string(data), string(content))
		}

		// Test with non-existent file (should panic)
		defer func() {
			if r := recover(); r == nil {
				t.Error("readFile should panic with non-existent file, but didn't")
			}
		}()
		readFile("nonexistent_file.txt")
	})
}
