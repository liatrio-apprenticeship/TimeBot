package main

import (
        "fmt"
        "log"
        "time"

	"golang.org/x/net/context"

        //"go.mongodb.org/mongo-driver/bson"
        
	"go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)

type Time struct {
    Timestamp time.Time
    In bool
}

func main() {
    /////////////////  MongoDB Stuff  ///////////////////////
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://database:27017")

    // Connect to MongoDB
    mongoclient, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = mongoclient.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }
    collection := mongoclient.Database("timesheets").Collection("daniel")
    /////////////////  End MongoDB Stuff  ///////////////////////


    time_in := Time{time.Now(), true}
    insertResult, err := collection.InsertOne(context.TODO(), time_in)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    fmt.Println("Clocked In")

}
