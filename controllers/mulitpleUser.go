package controllers

import (
	"net/http"
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	"github.com/gin-gonic/gin"
)
func MultipleUser(ctx *gin.Context){
	var users []models.User;
	err := ctx.ShouldBindJSON(&users);
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"invalid data","error":err.Error()})
		return;
	}
	var data []interface{};
	for _, user := range users{
		data = append(data,user);
	}
	err = mongoServices.InsertManyMethod("userCol", data);
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"invalid data","error":err.Error()})
		return;
	}
	ctx.JSON(http.StatusOK,gin.H{"msg":"Data inserted successfully"});
}