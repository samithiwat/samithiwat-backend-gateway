package model

type PaginationQueryParams struct {
	Limit int64 `query:"limit"`
	Page  int64 `query:"page"`
}
