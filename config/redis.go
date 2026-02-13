package config

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		fmt.Println(err.Error())
	}

	rdb := redis.NewClient(opt)

	return rdb
}
