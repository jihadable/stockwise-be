package services

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PasswordResetService interface {
	SendPasswordResetEmail(email string) error
	ResetPassword(token, newPassword string) error
}

type PasswordResetServiceImpl struct {
	*gorm.DB
	Redis *redis.Client
}

func (service *PasswordResetServiceImpl) SendPasswordResetEmail(email string) error {
	panic("")
}

func (service *PasswordResetServiceImpl) ResetPassword(token, newPassword string) error {
	panic("")
}

func NewPasswordResetService(config *config.Config) PasswordResetService {
	return &PasswordResetServiceImpl{DB: config.DB, Redis: config.Redis}
}
