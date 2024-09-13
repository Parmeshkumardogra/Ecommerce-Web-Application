package database

import (
	"context"
	"log"
	"time"
	"github.com/BMS/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() error {

	clientOptions := options.Client().ApplyURI(config.Config.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//connecting to mongodb
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error during mongo.Connect:", err)
		return err
	}

	//checking the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Println("Error during mongo.Connect:", err)
		return err
	}
	log.Println("Connected to MongoDB!")
	return nil

}

func DisconnectDB(){
	log.Println("Attempting to disconnecting from MongoDB!");
	if Client != nil{
		ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second);
		defer cancel();

		err := Client.Disconnect(ctx); 
		if err !=nil{
			log.Printf("Error disconnecting MongoDB: %v",err);
		}else{
			log.Println("Disconnected from MongoDB!")
		}
	}else{
		log.Println("MongoDB client is nil, cannot disconnect.");
	}
}