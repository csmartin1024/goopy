package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collections *CollectionGroup
var CTX = context.TODO()

type CollectionGroup struct {
	Goops *mongo.Collection
}

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func InitDatabase() {

	connectionInfo := getConnectionInfo()

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionInfo.databaseURI).SetRetryWrites(false)

	// Connect to MongoDB
	client, err := mongo.Connect(CTX, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database(connectionInfo.databaseName)
	goopCollection := db.Collection("goops")
	Collections = &CollectionGroup{Goops: goopCollection}
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	// log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

func getConnectionInfo() connectionInfo {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	dbURI := os.Getenv("MONGODB_URI")
	stringArr := strings.Split(dbURI, ":")
	dbName := stringArr[1][2:]
	return connectionInfo{databaseURI: dbURI, databaseName: dbName}
}

type connectionInfo struct {
	databaseURI  string
	databaseName string
}
