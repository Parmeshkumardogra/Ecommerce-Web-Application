package middleware

import (
	"net/http"

	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateForShortToken(ctx *gin.Context){
	shortToken := ctx.Request.Header.Get("Authorization");

	if shortToken == ""{
		ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{"msg":"you are not authorized to login"});
		return;
	}
	emailID, err := utils.VerifyShortToken(shortToken);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{"msg":"invalid token for authorization"});
		return;
	};

	ctx.Set("emailID",emailID);
	ctx.Next();
}

func AuthenticateForLongToken(ctx *gin.Context){
	longToken := ctx.Request.Header.Get("Authorization");
	if longToken == ""{
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg":"you are not authorized to login"})
		return;
	}
	emailID, err := utils.VerifyLongToken(longToken);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{"msg":"invalid token for authorization"});
		return;
	};	
	ctx.Set("emailID",emailID);
	ctx.Next();
}

