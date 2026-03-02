package request

type UserRequest struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Bio      *string `json:"bio"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Username string  `json:"username" validate:"required"`
	Bio      *string `json:"bio"`
}

type UpdatePasswordRequest struct {
	NewPasswrod string `json:"new_password" validate:"required"`
}
