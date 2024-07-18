package validation

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type testFieldLevel struct {
	field reflect.Value
	param string
}

// ExtractType implements validator.FieldLevel.
func (t testFieldLevel) ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool) {
	panic("unimplemented")
}

// Field implements validator.FieldLevel.
func (t testFieldLevel) Field() reflect.Value {
	return t.field
}

// FieldName implements validator.FieldLevel.
func (t testFieldLevel) FieldName() string {
	return t.field.String()
}

// GetStructFieldOK implements validator.FieldLevel.
func (t testFieldLevel) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	panic("unimplemented")
}

// GetStructFieldOK2 implements validator.FieldLevel.
func (t testFieldLevel) GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool) {
	panic("unimplemented")
}

// GetStructFieldOKAdvanced implements validator.FieldLevel.
func (t testFieldLevel) GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {
	panic("unimplemented")
}

// GetStructFieldOKAdvanced2 implements validator.FieldLevel.
func (t testFieldLevel) GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool) {
	panic("unimplemented")
}

// GetTag implements validator.FieldLevel.
func (t testFieldLevel) GetTag() string {
	panic("unimplemented")
}

// Param implements validator.FieldLevel.
func (t testFieldLevel) Param() string {
	panic("unimplemented")
}

// Parent implements validator.FieldLevel.
func (t testFieldLevel) Parent() reflect.Value {
	panic("unimplemented")
}

// StructFieldName implements validator.FieldLevel.
func (t testFieldLevel) StructFieldName() string {
	panic("unimplemented")
}

// Top implements validator.FieldLevel.
func (t testFieldLevel) Top() reflect.Value {
	panic("unimplemented")
}

func makeTestFieldLevel(field reflect.Value) validator.FieldLevel {
	return testFieldLevel{field: field, param: ""}
}

func TestValidateUUID(t *testing.T) {
	id := uuid.New()
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{id, id.String()},     // Valid UUID
		{uuid.Nil, nil},       // Nil UUID
		{"invalid-uuid", nil}, // Invalid format
		{123, nil},            // Wrong type
	}

	for _, tt := range tests {
		got := validateUUID(reflect.ValueOf(tt.input))
		if got != tt.expected {
			t.Errorf("validateUUID(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		amount string
		valid  bool
	}{
		{"100.50", true},
		{"0", true},
		{"-10", false},
		{"abc", false}, // Invalid format
	}

	for _, tt := range tests {
		fl := makeTestFieldLevel(reflect.ValueOf(tt.amount))
		if validateAmount(fl) != tt.valid {
			t.Errorf("validateAmount(%q) = %v, want %v", tt.amount, !tt.valid, tt.valid)
		}
	}
}

func TestValidateCurrency(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"USD", true},
		{"EUR", true},
		{"TRY", true},
		{"XYZ", false}, // Invalid code
	}

	for _, tt := range tests {
		fl := makeTestFieldLevel(reflect.ValueOf(tt.code))
		if validateCurrency(fl) != tt.valid {
			t.Errorf("validateCurrency(%q) = %v, want %v", tt.code, !tt.valid, tt.valid)
		}
	}
}

func TestValidateUserName(t *testing.T) {
	validUsernames := []string{"john_doe", "jane.doe123", "user_42"}
	invalidUsernames := []string{"", "john doe", "user@example.com", "user*name"}

	for _, username := range validUsernames {
		fl := makeTestFieldLevel(reflect.ValueOf(username))
		if !validateUserName(fl) {
			t.Errorf("validateUserName(%q) = false, want true", username)
		}
	}

	for _, username := range invalidUsernames {
		fl := makeTestFieldLevel(reflect.ValueOf(username))
		if validateUserName(fl) {
			t.Errorf("validateUserName(%q) = true, want false", username)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	validPasswords := []string{"Password123!", "S3cur3P@ss", "C0mpl3x_P@ssw0rd"}
	invalidPasswords := []string{"", "password", "12345678", "short"}

	for _, password := range validPasswords {
		fl := makeTestFieldLevel(reflect.ValueOf(password))
		if !validatePassword(fl) {
			t.Errorf("validatePassword(%q) = false, want true", password)
		}
	}

	for _, password := range invalidPasswords {
		fl := makeTestFieldLevel(reflect.ValueOf(password))
		if validatePassword(fl) {
			t.Errorf("validatePassword(%q) = true, want false", password)
		}
	}
}

func TestValidateSlug(t *testing.T) {
	validSlugs := []string{"my-blog-post", "product-123", "category-name"}
	invalidSlugs := []string{"", "My Blog Post", "product 123", "category/name"}

	for _, slug := range validSlugs {
		fl := makeTestFieldLevel(reflect.ValueOf(slug))
		if !validateSlug(fl) {
			t.Errorf("validateSlug(%q) = false, want true", slug)
		}
	}

	for _, slug := range invalidSlugs {
		fl := makeTestFieldLevel(reflect.ValueOf(slug))
		if validateSlug(fl) {
			t.Errorf("validateSlug(%q) = true, want false", slug)
		}
	}
}

func TestValidateLocale(t *testing.T) {
	validLocales := []string{"en-US", "tr-TR", "de-DE"}
	invalidLocales := []string{"", "e", "t"}

	for _, locale := range validLocales {
		fl := makeTestFieldLevel(reflect.ValueOf(locale))
		if !validateLocale(fl) {
			t.Errorf("validateLocale(%q) = false, want true", locale)
		}
	}

	for _, locale := range invalidLocales {
		fl := makeTestFieldLevel(reflect.ValueOf(locale))
		if validateLocale(fl) {
			t.Errorf("validateLocale(%q) = true, want false", locale)
		}
	}
}

func TestValidateGender(t *testing.T) {
	validGenders := []string{"male", "female"}
	invalidGenders := []string{"", "m", "f"}

	for _, gender := range validGenders {
		fl := makeTestFieldLevel(reflect.ValueOf(gender))
		if !validateGender(fl) {
			t.Errorf("validateGender(%q) = false, want true", gender)
		}
	}

	for _, gender := range invalidGenders {
		fl := makeTestFieldLevel(reflect.ValueOf(gender))
		if validateGender(fl) {
			t.Errorf("validateGender(%q) = true, want false", gender)
		}
	}
}

func TestValidatePhone(t *testing.T) {
	validPhones := []string{"+905551234567", "+12125551212", "+447700900123"}
	invalidPhones := []string{"", "05551234567", "12125551212", "7700900123"}

	for _, phone := range validPhones {
		fl := makeTestFieldLevel(reflect.ValueOf(phone))
		if !validatePhone(fl) {
			t.Errorf("validatePhone(%q) = false, want true", phone)
		}
	}

	for _, phone := range invalidPhones {
		fl := makeTestFieldLevel(reflect.ValueOf(phone))
		if validatePhone(fl) {
			t.Errorf("validatePhone(%q) = true, want false", phone)
		}
	}
}
