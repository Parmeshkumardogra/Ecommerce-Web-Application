package routes

import (
	"net/http"

	"github.com/BMS/controllers"
	accountcontroller "github.com/BMS/controllers/accountController"
	"github.com/BMS/controllers/loginController"

	"github.com/BMS/controllers/redisController"
	"github.com/BMS/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine{
	route := gin.Default();
	
	route.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{"msg":"PONG"});
	})

	route.POST("/signUp",loginController.Signup);

	route.POST("/login",loginController.LoginToSendOTP);

	//shortToken authentication
	authenticateOTPShortToken := route.Group("/otpTokenVerify");
	authenticateOTPShortToken.Use(middleware.AuthenticateForShortToken);
	authenticateOTPShortToken.POST("/verifyOTP", loginController.VerifyOTP);


	//long token authentication

	authenticateForLongToken := route.Group("/authorised");
	authenticateForLongToken.Use(middleware.AuthenticateForLongToken);
	// authenticateForLongToken.POST("/findOne",controllers.FindOne);
	authenticateForLongToken.POST("/createProfile",accountcontroller.CreateProfile);
	authenticateForLongToken.POST("/createAccount",accountcontroller.CreateSavingAccount);

	//db intereaction servcies

	
	route.POST("/insertMany",controllers.MultipleUser);
	
	route.POST("/findMany",controllers.FindManyUsers);

	route.POST("/findOne",controllers.FindOne);
	//temproary services
	route.POST("/getOTPfromRedis",redisController.FetchOTPFromRedis);

	route.POST("/verifyUserStatus",controllers.UserVerification);

	return route;
}

