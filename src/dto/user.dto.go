package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserDto struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	ImageUrl  string `json:"image_url" validate:"url"`
}

func ValidateUser(user UserDto) []*BadReqErrResponse {
	var errors []*BadReqErrResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := BadReqErrResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
