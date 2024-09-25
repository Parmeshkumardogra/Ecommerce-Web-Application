package models

type Config struct{
	MongoURI string `json:"mongoURI"`
	DBName string `json:"dbname"`
	CollectionName collectionCodes `json:"collectionName"`
}
type collectionCodes struct{
	MD01 string
	MD02 string
}
type FindUserRequest struct{
	Email string `json:"email" binding:"required"`
}
type UpdateUserStatus struct {
	UserID string `bson:"email" json:"userid" binding:"required"`
	IsVerified bool `bson:"isVerified" json:"verifiedStatus"`
	IsPriority string `bson:"isPriority" json:"priority"`
}
