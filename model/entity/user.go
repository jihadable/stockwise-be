package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id              string    `gorm:"column:id;primaryKey" json:"id"`
	Username        string    `gorm:"column:username;unique" json:"username"`
	Email           string    `gorm:"column:email;unique" json:"email"`
	Password        string    `gorm:"column:password" json:"password"`
	Bio             *string   `gorm:"column:bio" json:"bio"`
	IsEmailVerified bool      `gorm:"column:is_email_verified;default:false" json:"is_email_verified"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`

	Products Product `gorm:"foreignKey:UserId;references:Id" json:"products"`
}

func (model *User) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
