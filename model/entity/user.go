package entity

import "time"

type User struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	Username  string    `gorm:"column:username;unique" json:"username"`
	Email     string    `gorm:"column:email;unique" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Bio       string    `gorm:"column:bio" json:"bio"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Products Product `gorm:"foreignKey:UserId;references:Id" json:"products"`
}
