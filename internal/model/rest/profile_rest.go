package rest

type EditProfileRequest struct {
	Name string `json:"name" validate:"required"`
}
