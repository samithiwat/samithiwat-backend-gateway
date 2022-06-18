package dto

type BlogUserDto struct {
	UserId      uint32 `json:"user_id" validate:"required"`
	Description string `json:"description"`
}
