package request

type SendEmailVerificationRequest struct {
	Email string `json:"email" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}
