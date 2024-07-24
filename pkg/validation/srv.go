package validation

import (
	"context"
	"strings"

	"github.com/9ssi7/gopre/pkg/rescode"
	"github.com/9ssi7/gopre/pkg/state"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	ValidateStruct(ctx context.Context, sc interface{}) error
	ValidateMap(ctx context.Context, m map[string]interface{}, rules map[string]interface{}) error
}

type srv struct {
	validator *validator.Validate
	uni       *ut.UniversalTranslator
}

func New() Service {
	v := validator.New()
	v.RegisterCustomTypeFunc(validateUUID, uuid.UUID{})
	_ = v.RegisterValidation("username", validateUserName)
	_ = v.RegisterValidation("password", validatePassword)
	_ = v.RegisterValidation("locale", validateLocale)
	_ = v.RegisterValidation("slug", validateSlug)
	_ = v.RegisterValidation("gender", validateGender)
	_ = v.RegisterValidation("phone", validatePhone)
	_ = v.RegisterValidation("currency", validateCurrency)
	_ = v.RegisterValidation("amount", validateAmount)
	return &srv{validator: v, uni: ut.New(tr.New(), en.New())}
}

// ValidateStruct validates the given struct.
func (s *srv) ValidateStruct(ctx context.Context, sc interface{}) error {
	var errors []*ErrorResponse
	translator := s.getTranslator(ctx)
	err := s.validator.StructCtx(ctx, sc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			ns := s.mapStructNamespace(err.Namespace())
			if ns != "" {
				element.Namespace = ns
			}
			element.Field = err.Field()
			element.Value = err.Value()
			element.Message = err.Translate(translator)
			errors = append(errors, &element)
		}
	}
	if len(errors) > 0 {
		return rescode.ValidationFailed(nil).SetData(errors)
	}
	return nil
}

// ValidateMap validates the giveb struct.
func (s *srv) ValidateMap(ctx context.Context, m map[string]interface{}, rules map[string]interface{}) error {
	var errors []*ErrorResponse
	errMap := s.validator.ValidateMapCtx(ctx, m, rules)
	translator := s.getTranslator(ctx)
	for key, err := range errMap {
		var element ErrorResponse
		if _err, ok := err.(validator.ValidationErrors); ok {
			for _, err := range _err {
				element.Namespace = err.Namespace()
				element.Field = err.Field()
				if element.Field == "" {
					element.Field = key
				}
				element.Value = err.Value()
				element.Message = err.Translate(translator)
				errors = append(errors, &element)
			}
			continue
		}
	}
	if len(errors) > 0 {
		return rescode.ValidationFailed(nil).SetData(errors)
	}
	return nil
}

func (s *srv) getTranslator(ctx context.Context) ut.Translator {
	locale := state.GetLocale(ctx)
	translator, found := s.uni.GetTranslator(locale)
	if !found {
		translator = s.uni.GetFallback()
	}
	return translator
}

func (s *srv) mapStructNamespace(ns string) string {
	str := strings.Split(ns, ".")
	return strings.Join(str[1:], ".")
}
