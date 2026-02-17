package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmailVerification struct {
	Id       string    `grom:"column:id;primaryKey" json:"id"`
	Token    string    `gorm:"column:token" json:"token"`
	UserId   string    `gorm:"column:user_id" json:"user_id"`
	ExpireAt time.Time `gorm:"column:expire_at" json:"expire_at"`

	User *User `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

func (model *EmailVerification) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
