package dto

type OrganizationDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
