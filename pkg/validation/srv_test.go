package validation

import (
	"context"
	"strings"
	"testing"

	"github.com/9ssi7/gopre/pkg/rescode" // Update with your package path
	// Update with your package path
	"github.com/google/uuid"
)

// ... (existing validation functions: validateUUID, validateIban, etc.)

// Test struct
type TestStruct struct {
	ID       uuid.UUID `json:"id" validate:"uuid"`
	Name     string    `json:"name" validate:"required,username"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,password"`
	Locale   string    `json:"locale" validate:"required,locale"`
	Slug     string    `json:"slug" validate:"required,slug"`
	Gender   string    `json:"gender" validate:"required,gender"`
	Phone    string    `json:"phone" validate:"required,phone"`
}

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatal("New returned nil")
	}
}

func TestValidateStruct(t *testing.T) {
	s := New()

	// Valid struct
	validStruct := TestStruct{
		ID:       uuid.New(),
		Name:     "johndoe",
		Email:    "johndoe@example.com",
		Password: "ValidPassword123!",
		Locale:   "en-US",
		Slug:     "my-blog-post",
		Gender:   "male",
		Phone:    "+12125551212",
	}
	err := s.ValidateStruct(context.Background(), validStruct)
	if err != nil {
		t.Errorf("ValidateStruct() with valid struct returned an error: %v", err)
	}

	// Invalid struct with multiple errors
	invalidStruct := TestStruct{
		Name:     "",        // Missing required field
		Email:    "invalid", // Invalid email format
		Password: "short",   // Password too short
		Locale:   "invalid", // Invalid locale
		Slug:     "",        // Missing required field
		Gender:   "",        // Missing required field
		Phone:    "",        // Missing required field
	}
	err = s.ValidateStruct(context.Background(), invalidStruct)
	if err == nil {
		t.Fatal("ValidateStruct() with invalid struct did not return an error")
	}

	// Check if the returned error is of type rescode.Error
	_, ok := err.(*rescode.RC)
	if !ok {
		t.Errorf("ValidateStruct() did not return a rescode.Error")
	}

	// Check if the error messages are translated
	validationErrors, ok := err.(*rescode.RC).Data.([]*ErrorResponse)
	if !ok || len(validationErrors) == 0 {
		t.Fatal("ValidateStruct() did not return validation errors in data")
	}
	for _, e := range validationErrors {
		if e.Message == "" || !strings.Contains(e.Message, ": ") { // Check if the message contains a field name and a translated error
			t.Errorf("Validation error message is not translated: %s", e.Message)
		}
	}

	// Test with nil uuid
	invalidStruct.ID = uuid.Nil
	err = s.ValidateStruct(context.Background(), invalidStruct)
	if err == nil {
		t.Fatal("ValidateStruct() with nil UUID did not return an error")
	}
}

func TestValidateMap(t *testing.T) {
	s := New()

	validMap := map[string]interface{}{
		"name":     "johndoe",
		"email":    "johndoe@example.com",
		"password": "ValidPassword123!",
	}
	rules := map[string]interface{}{
		"name":     "required,username",
		"email":    "required,email",
		"password": "required,password",
	}

	err := s.ValidateMap(context.Background(), validMap, rules)
	if err != nil {
		t.Errorf("ValidateMap() with valid map returned an error: %v", err)
	}

	invalidMap := map[string]interface{}{
		"name":     "",        // Missing required field
		"email":    "invalid", // Invalid email format
		"password": "short",   // Password too short
	}

	err = s.ValidateMap(context.Background(), invalidMap, rules)
	if err == nil {
		t.Fatal("ValidateMap() with invalid map did not return an error")
	}

	// Check if the returned error is of type rescode.Error
	_, ok := err.(*rescode.RC)
	if !ok {
		t.Errorf("ValidateMap() did not return a rescode.Error")
	}
}
