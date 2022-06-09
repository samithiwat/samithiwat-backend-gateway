package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
)

type DtoValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

func (v *DtoValidator) Validate(in interface{}) []*dto.BadReqErrResponse {
	err := v.v.Struct(in)

	var errors []*dto.BadReqErrResponse
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			element := dto.BadReqErrResponse{
				Message:     e.Translate(v.trans),
				FailedField: e.StructField(),
				Tag:         e.Tag(),
				Value:       e.Value(),
			}

			errors = append(errors, &element)
		}
	}
	return errors
}

func NewValidator() (*DtoValidator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, errors.New("translator not found")
	}

	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, err
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("password", trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} is not strong enough (must be at lease 8 characters)", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 8
	})

	return &DtoValidator{
		v:     v,
		trans: trans,
	}, nil
}

func (v *DtoValidator) GetValidator() *validator.Validate {
	return v.v
}
