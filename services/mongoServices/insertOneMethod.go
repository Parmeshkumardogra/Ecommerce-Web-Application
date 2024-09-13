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
func InsertOneMethod(collectionName string, data interface{}) error{
	//start a new session
	session, err := database.Client.StartSession();
	if err != nil {
		log.Printf("failed to start session %v",err);
		return err;
	}
	defer session.EndSession(context.Background());

	//use context with time out to ensure that the transaction doesn't run infinite
	ctx, cancel := context.WithTimeout(context.Background(),10* time.Second)
	defer cancel();
	err = mongo.WithSession(ctx,session,  func(sc mongo.SessionContext) error {
		wc := writeconcern.Majority()
		err:= session.StartTransaction(options.Transaction().SetWriteConcern(wc));
		if err != nil {
			log.Printf("failed to start transaction: %v", err)
			return err;
		}
		//use collection of DB

		col := database.Client.Database(config.Config.DBName).Collection(collectionName);

		_, err = col.InsertOne(sc,data);
		if err != nil {
			log.Printf("error occurred during insert: %v", err);
			abortErr := session.AbortTransaction(sc);
			if abortErr != nil {
				log.Printf("failed to abort transaction: %v", abortErr);
				return abortErr;
			}
			return err;

		}

		//commit the transaction permenant
		err = session.CommitTransaction(sc);
		if err != nil{
			log.Printf("session commit error: %v", err)
			return err;
		}

		log.Printf("successfully inserted document into collection %s", collectionName)
		return nil;			
	})
	if err!= nil{
		log.Printf("transaction failed %v",err);
		return err;
	}
	return err;
}