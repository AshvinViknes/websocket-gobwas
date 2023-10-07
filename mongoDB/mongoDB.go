package mongoDB

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var chatHistoryCollection *mongo.Collection

func SetupMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	chatHistoryCollection = mongoClient.Database("chatapp").Collection("ChatHistory")
	log.Println("Connected to MongoDB successfully")
}

func SaveMessageToMongo(message string) error {
	doc := bson.M{
		"message":           message,
		"messageReceivedAt": time.Now(),
	}
	_, err := chatHistoryCollection.InsertOne(context.TODO(), doc)
	return err
}
