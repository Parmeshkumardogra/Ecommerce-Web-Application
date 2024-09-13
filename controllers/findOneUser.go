package controllers

import (
	"log"
	"net/http"

	"github.com/BMS/models"
	"github.com/BMS/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func FindOne(ctx *gin.Context){
	var request models.FindUserRequest;
	err := ctx.ShouldBindJSON(&request);
	log.Println("Requested Payload ",request.FirstName);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"invalid payload", "error":err.Error()});
		return
	}
	filterQueryPayload := bson.M{"firstName":request.FirstName};
	var result bson.M;
	result, err = services.FindOneMethod("userData",filterQueryPayload);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":"internal server error"});
		return;
	}
	if result == nil {
		ctx.JSON(http.StatusNotFound,gin.H{"msg":"data not found"});
		return;
	}
	ctx.JSON(http.StatusOK,gin.H{"msg":"Data found successfully","data":result});

}