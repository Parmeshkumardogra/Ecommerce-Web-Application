package services

import (
	"context"
	"fmt"
	"github.com/BMS/database"
	"github.com/go-redis/redis/v8"
)

func GetOTPFromRedis(userID string) (string, error) {
	validOTP, err := database.RedisClient.Get(context.Background(), userID).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("no OTP found for user ID: %s ", userID)
	} else if err != nil {
		return "", fmt.Errorf("error occured in fetching OTP %v ", err.Error())
	}
	return validOTP, nil
}
