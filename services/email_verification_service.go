package services

import (
	"time"

	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/helper"
	"github.com/jihadable/stockwise-be/model/entity"
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
	var emailTarget string
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		user := entity.User{}

		err := tx.Where("id = ?", userId).First(&user).Error
		if err != nil {
			return err
		}

		emailTarget = user.Email

		token, err := helper.GetToken()
		if err != nil {
			return err
		}
		expireAt := time.Now().Add(24 * time.Hour)

		emailVerification := entity.EmailVerification{
			Token:    token,
			UserId:   user.Id,
			ExpireAt: expireAt,
		}

		err = tx.Create(&emailVerification).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	emailVerificationLink := ""

	return helper.SendEmailVerification(emailTarget, emailVerificationLink)
}

func (service *EmailVerificationServiceImpl) VerifyEmail(token string) error {
	panic("")
}

func NewEmailVerificationService(config *config.Config) EmailVerificationService {
	return &EmailVerificationServiceImpl{DB: config.DB, Redis: config.Redis}
}
