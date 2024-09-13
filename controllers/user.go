package controllers

import (
	"github.com/BMS/models"
	"github.com/BMS/services/mongoServices"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "request is invalid", "error": err.Error()})
		return
	}
	err = mongoServices.InsertOneMethod("userData", user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "unable to insert data", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Data inserted successfully"})

}
