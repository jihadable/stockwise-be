package main

import (
	"fmt"
	"os"
	"strings"

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
		Username: "jihadable",
		Email:    "umarjihad@gmail.com",
		Password: hashedPassword,
	}

	err = db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("Fail to create user")
	}

	return nil
}

func seeder(db *gorm.DB) error {
	err := userSeeder(db)
	if err != nil {
		return fmt.Errorf("Fail to seed")
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
