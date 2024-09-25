package mongoServices

import (
	"context"
	"log"
	"time"

	"github.com/BMS/config"
	"github.com/BMS/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func InsertManyMethod(collectionName string, data []interface{}) error {
	session, err := database.Client.StartSession()
	if err != nil {
		log.Printf("error while creating sessions %v", err.Error())
		return err
	}
	defer session.EndSession(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongo.WithSession(ctx,session, func(sc mongo.SessionContext) error {
		wc := writeconcern.Majority()
		err = session.StartTransaction(options.Transaction().SetWriteConcern(wc))
		if err != nil {
			log.Printf("failed to start transaction: %v", err)
			return err
		}
		col := database.Client.Database(config.Config.DBName).Collection(collectionName)
		_, err := col.InsertMany(ctx, data)
		if err != nil {
			log.Printf("error occured at the time of inserting data %v", err.Error())
			err = session.AbortTransaction(context.Background());
			if err != nil{
				return err;
			}
			return err
		}
		err = session.CommitTransaction(ctx)
		if err != nil {
			log.Println("transaction is committed successfully")
			return err;
		}
		log.Printf("inserted %d documents into collection %s", len(data), collectionName)
		return nil;
	})
	if err != nil{
		log.Printf("error while mongo withSessoin operation %v",err.Error());
		return err;
	}
	return nil;
}
