package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserAddress struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Zip     string `bson:"zip" json:"zip"`
	Country string `bson:"country" json:"country"`
}
type RequestedUserProfilePayload struct {
	FirstName            string             `bson:"firstName" json:"firstName" binding:"required"`
	LastName             string             `bson:"lastName" json:"lastName" binding:"required"`
	Address              UserAddress        `bson:"address" json:"address" binding:"required"`
	DateOfBirth          string             ` json:"dateOfBirth" binding:"required"`
	AdharCardNo          string             `bson:"adharCardNo" json:"adharCardNo" binding:"required"`
	PanCardNo            string             `bson:"panCardNo" json:"panCardNo" binding:"required"`
}
type UserProfile struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName            string             `bson:"firstName" json:"firstName" binding:"required"`
	LastName             string             `bson:"lastName" json:"lastName" binding:"required"`
	Email                string             `bson:"email" json:"email"`
	Address              UserAddress        `bson:"address" json:"address" binding:"required"`
	DateOfBirth          time.Time          `bson:"dataOfBirth" json:"dateOfBirth" binding:"required"`
	AdharCardNo          string             `bson:"adharCardNo" json:"adharCardNo" binding:"required"`
	PanCardNo            string             `bson:"panCardNo" json:"panCardNo" binding:"required"`
	IsUserDetailsVerfied bool               `bson:"isUserDetailsVerified"`
	IsPriority           string             `bson:"isPriority"`
	Role                 string             `bson:"role"`
}
type UserCredentials struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email        string             `bson:"email" json:"email" binding:"required"`
	Phone        string             `bson:"phone" json:"phone" binding:"required"`
	PasswordHash string             `bson:"passwordHash" json:"password" binding:"required"`
	IsVerified   bool               `bson:"isVerified" json:"isVerified"`
}
