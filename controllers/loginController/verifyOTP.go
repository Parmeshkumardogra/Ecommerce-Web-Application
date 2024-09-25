package loginController

import (
	"log"
	"net/http"

	"github.com/BMS/services/redisServices"
	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
)

func VerifyOTP(ctx *gin.Context) {
	type Auth struct {
		Email string `json:"email" binding:"required"`
		OTP   string `json:"otp" binding:"required"`
	}
	var auth Auth
	err := ctx.ShouldBindJSON(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid payload"})
		return
	}
	var userOTP string
	userOTP,err = redisServices.GetOTPFromRedis(auth.Email);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if auth.OTP != userOTP {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Incorrect OTP"})
		return
	}
	log.Println("Your OTP is ",userOTP," for userID ",auth.Email);
	var jwtLongToken string
	jwtLongToken, err = utils.GenerateLongToken(auth.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "login successfull", "jwtLongToken": jwtLongToken})
}
