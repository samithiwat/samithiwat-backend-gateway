package dto

type UserDto struct {
	Firstname   string `json:"firstname" validate:"required"`
	Lastname    string `json:"lastname" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"url"`
}
