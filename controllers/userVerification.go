package controllers

import (
	"github.com/BMS/config"
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func UserVerification(ctx *gin.Context) {
	var user models.UpdateUserStatus
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "false", "error": err.Error()})
		return
	}
	//validate priority skip
	log.Println(&user)
	var updateData bson.M
	if !user.IsVerified && (user.IsPriority == "" || user.IsPriority == "LW") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "false", "error": "invalid request"})
		return
	} else if user.IsVerified && user.IsPriority == "" {
		updateData = bson.M{"$set": bson.M{"isVerified": user.IsVerified}}

	} else if !user.IsVerified && user.IsPriority != "" {
		updateData = bson.M{"$set": bson.M{"isPriority": user.IsPriority}}
	} 
	if user.IsVerified && user.IsPriority != ""  {
		updateData = bson.M{"$set": bson.M{"isPriority": user.IsPriority, "isVerified": user.IsVerified}}
	}
	filter := bson.M{"email": user.UserID}
	err = mongoServices.FindOneAndUpdate(config.Config.CollectionName.MD01, filter, updateData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "false", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "msg": "data updated successfully"})
}
