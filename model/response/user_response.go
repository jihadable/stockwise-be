package response

type UserResponse struct {
	Id              string  `json:"id"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	Bio             *string `json:"bio"`
	IsEmailVerified bool    `json:"is_email_verified"`
}
