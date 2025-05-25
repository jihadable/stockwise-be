package request

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Category    string  `json:"category" validate:"category"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,min=1"`
	Description string  `json:"description" validate:"required"`
	UserId      string  `json:"user_id"`
}
