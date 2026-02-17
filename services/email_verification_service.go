package services

import (
	"errors"
	"fmt"
	"os"
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
	var emailTarget, token string
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

	emailVerificationLink := fmt.Sprintf("%s/verify-email/%s", os.Getenv("WEB_ENDPOINT"), token)

	return helper.SendEmailVerification(emailTarget, emailVerificationLink)
}

func (service *EmailVerificationServiceImpl) VerifyEmail(token string) error {
	return service.DB.Transaction(func(tx *gorm.DB) error {
		emailVerification := entity.EmailVerification{}

		err := tx.Where("token = ? AND expire_at > ?", token, time.Now()).First(&emailVerification).Error
		if err != nil {
			return err
		}

		result := tx.Model(&entity.User{}).Where("id = ? AND is_email_verified = ?", emailVerification.UserId, false).Update("is_email_verified", true)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("User already verified")
		}

		err = tx.Delete(&emailVerification).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func NewEmailVerificationService(config *config.Config) EmailVerificationService {
	return &EmailVerificationServiceImpl{DB: config.DB, Redis: config.Redis}
}
