package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type PaginationQueryParams struct {
	Limit int64 `query:"limit"`
	Page  int64 `query:"page"`
}

type ResponseErr struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type BadReqErrResponse struct {
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}
