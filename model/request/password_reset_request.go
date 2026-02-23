package request

type SendPasswordResetEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
