package loginController

import (
	"github.com/BMS/config"
	"github.com/BMS/services/mongoServices"
	"github.com/BMS/services/redisServices"
	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "log"
	"net/http"
	"time"
)

func LoginToSendOTP(ctx *gin.Context) {
	type Credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var loginCredential Credentials

	//Binding requested payload & also check isValid payload or not
	err := ctx.ShouldBindJSON(&loginCredential)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid payload passed"})
		return
	}

	//Creating filter query payload to pass as a parameter in findone service
	filterQueryPayload := bson.M{"email": loginCredential.Email}
	var userDetails bson.M
	userDetails, err = mongoServices.FindOneMethod(config.Config.CollectionName.MD01, filterQueryPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "internal server error"})
		return
	}

	passwordHash, ok := userDetails["passwordHash"].(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "password not found"})
		return
	}
	// log.Print(password);
	
	//Compare the requested password & saved password
	ok = utils.CheckPassword(loginCredential.Password, passwordHash)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "incorrect passowrd"})
		return
	}

	//Generate OTP
	otp := utils.GenerateOTP()
	
	redisServices.SetOTPInRedis(loginCredential.Email, otp, 5*time.Minute)

	var jwtShortToken string
	jwtShortToken, err = utils.GenerateShortToken(loginCredential.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "OTP sent successfully", "jwtShortToken": jwtShortToken})
}
