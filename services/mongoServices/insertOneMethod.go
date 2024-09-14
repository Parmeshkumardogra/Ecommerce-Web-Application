package mongoServices

import (
	"context"
	"log"
	"github.com/BMS/config"
	"github.com/BMS/database"

)
func InsertOneMethod(collectionName string, data interface{}) error{
		//use collection of DB
		col := database.Client.Database(config.Config.DBName).Collection(collectionName);
		_, err := col.InsertOne(context.Background(),data);
		if err != nil {
			log.Printf("error occurred during insert: %v", err);
			return err;
		}
		log.Printf("successfully inserted document into collection %s", collectionName)
		return nil;			
}