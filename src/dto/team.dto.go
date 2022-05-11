package dto

type TeamDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
