package entity

import "time"

type Product struct {
	Id          string    `gorm:"column:id;primaryKey" json:"id"`
	Slug        string    `gorm:"column:slug;unique" json:"slug"`
	Name        string    `gorm:"column:name" json:"name"`
	Price       float32   `gorm:"column:price" json:"price"`
	Quantity    int       `gorm:"column:quantity" json:"quantity"`
	Image       *string   `gorm:"column:image" json:"image"`
	Description string    `gorm:"column:description" json:"description"`
	UserId      string    `gorm:"column:user_id" json:"user_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`

	User *User `gorm:"foreignKey:UserId;references:Id" json:"user"`
}
