package accountcontroller

import (
	"github.com/BMS/config"
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateProfile(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "missing or invalid token"})
		return
	}
	email, err := utils.VerifyLongToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}
	var requestedPayload models.RequestedUserProfilePayload
	err = ctx.ShouldBindJSON(&requestedPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid payload received", "error": err.Error()})
		return
	}
	priorityList := [3]string{"LW", "MD", "HG"}
	var dateofBirthTimeFormat time.Time;
	dateofBirthTimeFormat, err = time.Parse("2006-01-02", requestedPayload.DateOfBirth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid date received", "error": err.Error()})
		return
	}
	var userProfile models.UserProfile
	userProfile.FirstName = requestedPayload.FirstName
	userProfile.LastName = requestedPayload.LastName
	userProfile.Address = requestedPayload.Address
	userProfile.AdharCardNo = requestedPayload.AdharCardNo
	userProfile.PanCardNo=requestedPayload.PanCardNo
	userProfile.DateOfBirth = dateofBirthTimeFormat
	userProfile.Email = email
	userProfile.ID = primitive.NewObjectID()
	userProfile.IsUserDetailsVerfied = false
	userProfile.IsPriority = priorityList[0]
	userProfile.Role="customer"

	err = mongoServices.InsertOneMethod(config.Config.CollectionName.MD02, userProfile)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error", "error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "user details saved we will verify the details and updated you in 24-48 hours."})

}
