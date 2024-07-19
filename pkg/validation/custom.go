package validation

import (
	"reflect"
	"regexp"

	"github.com/9ssi7/gopre/pkg/currency"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func validateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		if valuer.String() == uuid.Nil.String() {
			return nil
		}
		return valuer.String()
	}
	return nil
}

func validateAmount(fl validator.FieldLevel) bool {
	d, err := decimal.NewFromString(fl.Field().String())
	if err != nil {
		return false
	}
	return d.GreaterThanOrEqual(decimal.Zero)
}

func validateCurrency(fl validator.FieldLevel) bool {
	return currency.IsValid(fl.Field().String())
}

func validateUserName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(userNameRegexp, fl.Field().String())
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(passwordRegexp, fl.Field().String())
	return matched
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(slugRegexp, fl.Field().String())
	return matched
}

func validateLocale(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(localeRegexp, fl.Field().String())
	return matched
}

func validateGender(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(genderRegexp, fl.Field().String())
	return matched
}

func validatePhone(f1 validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(phoneWithCountryCodeRegexp, f1.Field().String())
	return matched
}
