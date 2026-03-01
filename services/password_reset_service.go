package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/helper"
	"github.com/jihadable/stockwise-be/helper/mailer"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
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
	var token string

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		user := entity.User{}

		err := tx.Where("email = ?", email).First(&user).Error
		if err != nil {
			return err
		}

		token, err = helper.GetToken()
		if err != nil {
			return err
		}
		expireAt := time.Now().Add(24 * time.Hour)

		passwordReset := entity.PasswordReset{
			Token:    token,
			UserId:   user.Id,
			ExpireAt: expireAt,
		}
		err = tx.Create(&passwordReset).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	passwordResetLink := fmt.Sprintf("%s/reset-password/%s", os.Getenv("WEB_ENDPOINT"), token)

	return mailer.SendPasswordReset(email, passwordResetLink)
}

func (service *PasswordResetServiceImpl) ResetPassword(token, newPassword string) error {
	return service.DB.Transaction(func(tx *gorm.DB) error {
		passwordReset := entity.PasswordReset{}

		err := tx.Where("token = ?", token).First(&passwordReset).Error
		if err != nil {
			return err
		}

		user := entity.User{}

		err = tx.Where("id = ?", passwordReset.UserId).First(&user).Error
		if err != nil {
			return err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newPassword))
		if err == nil {
			return errors.New("New password can not be same as old password")
		}

		hashedNewPassword, err := helper.HashPassword(newPassword)
		if err != nil {
			return err
		}

		user.Password = hashedNewPassword

		err = tx.Updates(&user).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func NewPasswordResetService(config *config.Config) PasswordResetService {
	return &PasswordResetServiceImpl{DB: config.DB, Redis: config.Redis}
}
