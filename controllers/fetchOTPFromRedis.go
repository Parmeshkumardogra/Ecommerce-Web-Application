package controllers

import (
	"net/http"

	"github.com/BMS/services/redisServices"
	"github.com/gin-gonic/gin"
)

func FetchOTPFromRedis(ctx *gin.Context){
	type UserID struct{
		Email string `json:"email" binding:"required"`
	}
	var userID UserID
	err := ctx.ShouldBindJSON(&userID);
	if err !=nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error":err.Error()});
		return;
	}
	var userOTP string;
	userOTP,err = redisServices.GetOTPFromRedis(userID.Email);
	if err !=nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()});
		return;
	}
	ctx.JSON(http.StatusOK,gin.H{"otp":userOTP,"userID":userID.Email});
}