package controllers

import (
	"log"
	"net/http"
	"github.com/BMS/models"
	"github.com/BMS/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindManyUsers(ctx *gin.Context) {
	var request models.FindUserRequest
	err := ctx.ShouldBindJSON(&request)
	log.Println("Requested payload ", request)
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid payload", "error": err.Error()})
		return
	}
	var results []bson.M
	
	filterQueryPayload := bson.M{"firstName": request.FirstName}
	results, err = services.FindMethod("userData", filterQueryPayload)
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "error occured in finding data", "error": err.Error()})
		return
	}
	
	if len(results) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "Data not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"msg": "Data found successfully", "data": results})

}
