package main

import (
	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migrator(db *gorm.DB) error {
	err := db.Migrator().DropTable(&entity.User{}, &entity.Product{}, &entity.EmailVerification{})
	if err != nil {
		return err
	}

	return db.Migrator().CreateTable(&entity.User{}, &entity.Product{}, &entity.EmailVerification{})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db := config.DB()

	err = migrator(db)
	if err != nil {
		panic(err)
	}
}
