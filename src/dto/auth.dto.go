package dto

type Register struct {
	Email       string `json:"email" validate:"email" example:"admin@samithiwat.dev"`
	Password    string `json:"password" validate:"password" example:"password"`
	Firstname   string `json:"firstname" validate:"required" example:"Samithiwat"`
	Lastname    string `json:"lastname" validate:"required" example:"Boonchai"`
	DisplayName string `json:"display_name" validate:"required" example:"Smithy"`
	ImageUrl    string `json:"image_url" example:"https://storage.googleapis.com/samithiwat-bucket/about-me-protrait.png"`
}

type Login struct {
	Email    string `json:"email" validate:"email" example:"admin@samithiwat.dev"`
	Password string `json:"password" validate:"required" example:"password"`
}

type ChangePassword struct {
	UserId      uint32 `json:"user_id" validate:"required" example:"1"`
	OldPassword string `json:"old_password" validate:"required" example:"password"`
	NewPassword string `json:"new_password" validate:"password" example:"new_password"`
}

type Validate struct {
	Token string `json:"token" validate:"jwt"`
}

type RedeemNewToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
