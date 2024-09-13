package services

import (
	"context"
	"log"

	"github.com/BMS/config"
	"github.com/BMS/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func FindOneMethod(collectionName string, filter interface{}) (bson.M, error) {
	col := database.Client.Database(config.Config.DBName).Collection(collectionName);
	var response bson.M;
	err :=  col.FindOne(context.Background(),filter).Decode(&response)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return nil,nil;
		}
		log.Printf("Error finding document: %v", err)
		return nil, err;
	}
	return response, nil;
}