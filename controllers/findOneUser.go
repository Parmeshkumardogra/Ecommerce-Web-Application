package controllers

import (
	"log"
	"net/http"
	"github.com/BMS/config"
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func FindOne(ctx *gin.Context){
	var request models.FindUserRequest;
	err := ctx.ShouldBindJSON(&request);
	log.Println("Requested Payload ",request.Email);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"invalid payload", "error":err.Error()});
		return
	}
	filterQueryPayload := bson.M{"email":request.Email};
	var result bson.M;
	result, err = mongoServices.FindOneMethod(config.Config.CollectionName.MD01,filterQueryPayload);
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