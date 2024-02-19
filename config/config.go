package config

import (
	"context"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "testApi"
const User = "users"
const Role = "role"

var UserCollection *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	client := options.Client().ApplyURI(os.Getenv(os.Getenv("RUN_MODE")+"_DB_URL"))

	connection, err := mongo.Connect(context.TODO(), client)
	if err != nil {
		log.Fatal(err)
		color.Blue("DATABASE CONNECTION FAIL" + os.Getenv(os.Getenv("RUN_MODE")+"_DB_URL"))

	}
	color.Blue("DATABASE CONNECTION SUCCESSFUL  :: " + os.Getenv(os.Getenv("RUN_MODE")+"_DB_URL"))
	UserCollection = connection.Database(dbName).Collection(User)
	
}

