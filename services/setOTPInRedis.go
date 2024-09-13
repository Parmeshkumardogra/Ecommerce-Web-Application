package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BMS/database"
)

func SetOTPInRedis(userID, otp string, otpExpiryTime time.Duration) error {

	err := database.RedisClient.Set(context.Background(), userID, otp, otpExpiryTime).Err()
	if err != nil {
		log.Printf("failed to set OTP in redis %v", err.Error())
		return fmt.Errorf("failed to set OTP in redis %v", err.Error())
	}
	return nil

}
