package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

)
type Address struct{
	Street string `bson:"street" json:"street"`
	City string `bson:"city" json:"city"`
	State string `bson:"state" json:"state"`
	Zip string `bson:"zip" json:"zip"`
	Country string `bson:"country" json:"country"`
}

type Account struct{
	AccountID primitive.ObjectID `bson:"accountID" json:"accountID"`
	AccountType string `bson:"accountType" json:"accountType"`
	Balance float64 `bson:"balance" json:"balance"`
	OpenedOn time.Time `bson:"openedOn" json:"openedOn"`
	Status string `bson:"status" json:"status"`
}

type Customer struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	// CustomerID string `bson:"customerID" json:"customerID"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Email string `bson:"email" json:"email" binding:"required"`
	Phone string `bson:"phone" json:"phone"`
	Address Address `bson:"address" json:"address"`
	DateOfBirth time.Time `bson:"dataOfBirth" json:"dateOfBirth"`
	PasswordHash string `bson:"passwordHash" json:"password" binding:"required"`
	// PasswordSalt string `bson:"passwordSalt" json:"passwordSalt"`
	Accounts []Account `bson:"accounts" json:"accounts"`
	// VerificationStatus string `bson:"verificationStatus"`
	IsVerified bool `bson:"isVerified"`
	IsPriority string `bson:"isPriority"`
}