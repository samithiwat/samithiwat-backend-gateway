package model

type PaginationQueryParams struct {
	Limit int64 `query:"limit"`
	Page  int64 `query:"page"`
}

type ResponseErr struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
