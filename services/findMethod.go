package services

import (
	"context"
	"github.com/BMS/database"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/BMS/config"
	"log"
)

func FindMethod(collectionName string, filter interface{}) ([]bson.M, error) {
	
	log.Println("query for filter ",filter);
	col := database.Client.Database(config.Config.DBName).Collection(collectionName)
	var results []bson.M
	
	cursor, err := col.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	
	defer cursor.Close(context.Background())
	
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Printf("error occured in decoding result %v", err)
			return nil, err
		}
		results = append(results,result);
	}
	
	if cursor.Err() != nil{
		log.Printf("cursor error %v",err);
		return nil, err;
	}
	return results,nil;
}
