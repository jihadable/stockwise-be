package services

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type EmailVerificationService interface {
	SendEmailVerification(userId string) error
	VerifyEmail(token string) error
}

type EmailVerificationServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *EmailVerificationServiceImpl) SendEmailVerification(userId string) error {
	panic("")
}

func (service *EmailVerificationServiceImpl) VerifyEmail(token string) error {
	panic("")
}

func NewEmailVerificationService(config *config.Config) EmailVerificationService {
	return &EmailVerificationServiceImpl{DB: config.DB, Redis: config.Redis}
}
