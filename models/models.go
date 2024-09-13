package models

import "time"

type Config struct{
	MongoURI string `json:"mongoURI"`
	DBName string `json:"dbname"`
	CollectionName collectionCodes `json:"collectionName"`
}
type collectionCodes struct{
	MD01 string
}
type User struct{
	FirstName string `bson:"firstName" json:"firstName" binding:"required"`
	LastName string  `bson:"lastname" json:"lastName" binding:"required"`
}
type FindUserRequest struct{
	FirstName string `json:"firstName" binding:"omitempty"`
}
type MappedData struct{
	VirtualNumber string `json:"virtualNumber" binding:"required"`
	AgentNumber string `json:"agentNumber" binding:"required"`
	CustomerNumber string `json:"customerNumber" binding:"required"`
}
type MappedDataArray struct{
	Data []MappedData `json:"data" binding:"required"`
	TTL time.Duration `json:"ttl"`
}
type RemoveMapping struct{
	VirtualNumber string `json:"virtualNumber" binding:"required"`
	CustomerNumber string `json:"customerNumber" binding:"required"`
}
type RemoveMappedArray struct{
	Data []RemoveMapping `json:"data" binding:"required"`
}