package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	DB    *gorm.DB
	Redis *redis.Client
}
