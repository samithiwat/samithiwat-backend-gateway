package dto

type OrganizationDto struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Description string `json:"description"`
}
