package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/helper"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func truncateAllTable(db *gorm.DB) error {
	tables := []string{
		"users",
		"products",
		"email_verifications",
	}

	query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", strings.Join(tables, ", "))

	if err := db.Exec(query).Error; err != nil {
		return fmt.Errorf("failed to truncate tables: %w", err)
	}

	return nil
}

func userSeeder(db *gorm.DB) error {
	hashedPassword, err := helper.HashPassword(os.Getenv("PRIVATE_PASSWORD"))
	if err != nil {
		return fmt.Errorf("Fail to hash password")
	}

	user := entity.User{
		Username: "jihadumar",
		Email:    "jihadumar1021@gmail.com",
		Password: hashedPassword,
	}

	err = db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("Fail to create user")
	}

	return nil
}

func emailVerificationSeeder(db *gorm.DB) error {
	user := entity.User{}

	err := db.Where("email = ?", "jihadumar1021@gmail.com").First(&user).Error
	if err != nil {
		return err
	}

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

	err = db.Create(&emailVerification).Error
	if err != nil {
		return err
	}

	fmt.Println("Email verification token:", emailVerification.Token)

	return nil
}

func passwordResetSeeder(db *gorm.DB) error {
	user := entity.User{}

	err := db.Where("email = ?", "jihadumar1021@gmail.com").First(&user).Error
	if err != nil {
		return err
	}

	token, err := helper.GetToken()
	if err != nil {
		return err
	}
	expireAt := time.Now().Add(24 * time.Hour)

	passwordReset := entity.PasswordReset{
		Token:    token,
		ExpireAt: expireAt,
		UserId:   user.Id,
	}
	err = db.Create(&passwordReset).Error
	if err != nil {
		return err
	}

	fmt.Println("Password reset token:", passwordReset.Token)

	return nil
}

func seeder(db *gorm.DB) error {
	err := userSeeder(db)
	if err != nil {
		return fmt.Errorf("Fail to seed users")
	}

	err = emailVerificationSeeder(db)
	if err != nil {
		return fmt.Errorf("Fail to seed email verifications")
	}

	err = passwordResetSeeder(db)
	if err != nil {
		return fmt.Errorf("Fail to seed password resets")
	}

	return nil
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	db := config.DB()

	err = truncateAllTable(db)
	if err != nil {
		panic(err)
	}

	err = seeder(db)
	if err != nil {
		panic(err)
	}
}
