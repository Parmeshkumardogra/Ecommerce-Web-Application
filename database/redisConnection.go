package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func ConnectRedis() (error){
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
		return err;
	}
	log.Println("Connected to Redis!")
	return nil;
}
func DisconnectRedis() {
	log.Println("Attempting to disconnect from Redis")
	err := RedisClient.Close()
	if err != nil {
		log.Printf("Error disconnecting from Redis: %v", err)
	} else {
		log.Println("Disconnected from Redis!")
	}
}
