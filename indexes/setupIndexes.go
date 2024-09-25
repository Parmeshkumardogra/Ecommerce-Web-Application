package indexes

import (
	"context"
	"log"

	"github.com/BMS/config"
	"github.com/BMS/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetUpIndexes() error{
	col := database.Client.Database(config.Config.DBName).Collection(config.Config.CollectionName.MD01)
	_, err := col.Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys: bson.D{{Key:"email", Value:1}},
		Options: options.Index().SetUnique(true),
	});
	if err != nil {
		log.Fatalf("failed to create unique indexes %v",err.Error());
		return err;
	}
	_, err = col.Indexes().CreateOne(context.Background(),mongo.IndexModel{
		Keys:bson.D{{Key:"isVerified",Value:1}},
		Options: options.Index().SetPartialFilterExpression(bson.D{{Key:"isVerified",Value:false}}),
	})
	if err != nil {
		log.Fatalf("failed to create partial indexes %v",err.Error());
		return err;
	}
	log.Println("MongoDB Indexes ensured successfully!")
	return nil;
}