package mongoServices

import (
	"context"
	"log"
	"github.com/BMS/database"
	"github.com/BMS/config"
)
func InsertManyMethod(collectionName string, data []interface{}) error{
	col := database.Client.Database(config.Config.DBName).Collection(collectionName);
	_,err :=col.InsertMany(context.Background(),data);
	if err != nil {
		log.Printf("error occured at the time of inserting data %s",err.Error());
		return err
	}
	log.Printf("inserted %d documents into collection %s",len(data),collectionName);
	return nil;
}