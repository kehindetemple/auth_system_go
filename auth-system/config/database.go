package config

import (
	"log"
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)
var DB *mongo.Database
var UserCollection *mongo.Collection

func ConnectDB(){
	err := godotenv.Load()

	if err != nil{
		log.Fatal(err)
	}
	client, err := mongo.Connect(
			options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)
	if err != nil {
    	log.Fatal(err)
	}
	ctx := context.Background()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully")

	DB = client.Database("auth-system")
	UserCollection = DB.Collection("users")


}