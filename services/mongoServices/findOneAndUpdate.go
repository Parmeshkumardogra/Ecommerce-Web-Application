package mongoServices

import (
	"context"
	"log"
	"github.com/BMS/config"
	"github.com/BMS/database"
)

func FindOneAndUpdate(collectionName string, filter, data interface{}) error {
	col := database.Client.Database(config.Config.DBName).Collection(collectionName)
	res := col.FindOneAndUpdate(context.Background(), filter, data)
	if res.Err() != nil {
		log.Printf("error while find and update document %v", res.Err())
		return res.Err()
	}
	log.Printf("data updated succesfully in the %s collection", collectionName)
	return nil
}
