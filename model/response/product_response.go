package response

import "time"

type ProductResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity"`
	Image       *string   `json:"image"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
