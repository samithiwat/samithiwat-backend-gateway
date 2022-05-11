package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
)

var Validate = validator.New()

func Format(err validator.ValidationErrors) []*dto.BadReqErrResponse {
	var errors []*dto.BadReqErrResponse
	if err != nil {
		for _, err := range err {
			element := dto.BadReqErrResponse{
				FailedField: err.StructField(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			}

			errors = append(errors, &element)
		}
	}
	return errors
}
