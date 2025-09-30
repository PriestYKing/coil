package rediscache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	// Redis
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	//	dbname := os.Getenv("REDIS_DB")
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	// Ping to test connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis!")
	return rdb
}

func StoreJWTToken(rdb *redis.Client, ctx context.Context, email string, token string, expiresIn time.Duration) error {
	key := "auth:token" + email
	return rdb.Set(ctx, key, token, expiresIn).Err()
}

func DeleteJWTToken(rdb *redis.Client, ctx context.Context, email string) error {
	key := "auth:token" + email
	return rdb.Del(ctx, key).Err()
}
