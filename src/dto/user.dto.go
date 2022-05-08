package dto

type CreateUserDto struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ImageUrl  string `json:"image_url"`
}

type UpdateUserDto struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ImageUrl  string `json:"image_url"`
}
