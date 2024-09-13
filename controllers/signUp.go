package controllers

import (
	"net/http"

	"github.com/BMS/config"
	"github.com/BMS/models"
	"github.com/BMS/services"
	"github.com/BMS/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(ctx *gin.Context) {
	var customer models.Customer
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Invalid Payload", "error":err.Error()})
		return
	}
	customer.ID = primitive.NewObjectID()
	customer.CustomerID = primitive.NewObjectID().Hex()
	var hashedPassword string
	hashedPassword, err = utils.HashPassword(customer.PasswordHash)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customer.PasswordHash = hashedPassword
	err = services.InsertOneMethod(config.Config.CollectionName.MD01, customer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failing to insert data", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Customer onboarded successfully"})
}
