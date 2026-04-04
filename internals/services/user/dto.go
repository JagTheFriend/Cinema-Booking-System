package user

type SignUpRequest struct {
	ID string `json:"id" validate:"required,min=3"`
}

type LoginRequest struct {
	ID string `json:"id" validate:"required"`
}

type UpdateUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type UserResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
