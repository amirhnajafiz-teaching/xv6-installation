package redis

import "github.com/go-redis/redis/v9"

// NewConnection
// opens a connection to redis database.
func NewConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
