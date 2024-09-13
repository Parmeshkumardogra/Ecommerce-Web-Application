package routes

import (
	"net/http"

	"github.com/BMS/controllers"
	"github.com/BMS/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine{
	route := gin.Default();
	
	route.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{"msg":"PONG"});
	})

	route.POST("/signUp",controllers.Signup);

	route.POST("/generateShortToken",controllers.LoginToSendOTP);

	//shortToken authentication
	authenticateOTPShortToken := route.Group("/otpTokenVerify");
	authenticateOTPShortToken.Use(middleware.AuthenticateForShortToken);
	authenticateOTPShortToken.POST("/verifyOTP", controllers.VerifyOTP);


	//long token authentication

	authenticateForLongToken := route.Group("/longToken");
	authenticateForLongToken.Use(middleware.AuthenticateForLongToken);
	// authenticateForLongToken.POST("/findOne",controllers.FindOne);

	//db intereaction servcies

	route.POST("/insertOne",controllers.User)
	
	route.POST("/insertMany",controllers.MultipleUser);
	
	route.POST("/findMany",controllers.FindManyUsers);

	route.POST("/findOne",controllers.FindOne);
	//temproary services
	route.POST("/getOTPfromRedis",controllers.FetchOTPFromRedis);

	return route;
}

