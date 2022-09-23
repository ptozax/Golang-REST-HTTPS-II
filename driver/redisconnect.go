package driver

import (
	"github.com/go-redis/redis"
)

func RedisConnect() (*redis.Client, error) {

	//addr := os.Getenv("ADDR_REDIS")
	//pass := os.Getenv("PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb, nil
}
func RedisConnect_ID(id int) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       id,
	})

	return rdb, nil
}
