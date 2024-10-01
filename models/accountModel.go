package models
type RequestedCreateAccount struct{
	AccountTypeName string `json:"accountType" binding:"required"`
}