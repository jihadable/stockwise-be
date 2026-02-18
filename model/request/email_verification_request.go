package request

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}
