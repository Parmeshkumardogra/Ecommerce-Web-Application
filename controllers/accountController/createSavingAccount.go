package accountcontroller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/BMS/config"
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	mysqlservices "github.com/BMS/services/mysqlServices"
	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSavingAccount(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	email, err := utils.VerifyLongToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized", "error": "invalid token for authorization"})
		return
	}
	var accountType models.RequestedCreateAccount
	err = ctx.ShouldBindJSON(&accountType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid requested payload", "error": err.Error()})
		return
	}
	if accountType.AccountTypeName != "saving" && accountType.AccountTypeName != "credit" && accountType.AccountTypeName != "current" && accountType.AccountTypeName != "fixed" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid requested payload"})
		return
	}
	filterQuery := bson.M{"email":email}
	response, err := mongoServices.FindOneMethod(config.Config.CollectionName.MD02,filterQuery);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"status":false,"message":"internal server error"});
		return;
	}
	if response == nil {
		ctx.JSON(http.StatusNotFound,gin.H{"status":true,"message":"no data found"});
		return;
	}
	isUserDetailsVerified, ok := response["isUserDetailsVerified"];
	if !ok {
		ctx.JSON(http.StatusNotFound,gin.H{"status":true,"message":"no data found"});
		return;
	}
	if !isUserDetailsVerified.(bool){
		ctx.JSON(http.StatusFound, gin.H{"status":true,"message":"Your address verfication is still under process, kindly wait we will update you then please try again."});
		return;
	}
	query := `select accountTypeID from masterAccountType where accountTypeName = ?`;
	accTypeName := "saving";

	sqlRes := mysqlservices.SelectOne(query, accTypeName);
	var accountTypeID int;
	err = sqlRes.Scan(&accountTypeID);
	if err != nil {
		if err == sql.ErrNoRows{
			log.Println("no record found");
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status":false,"message":"internal server error" ,"error":err.Error()})
		return;
	}
	ctx.JSON(http.StatusOK,gin.H{"result":accountTypeID});
	
	type accountDetails struct {
		accountTypeID int
		accountID     string
		accountNumber string
		interestRate  float32
		createdBy     string
	}

}
