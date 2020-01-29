package main

import (
    "context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
    uid string
    sid string
}

func main() {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

    collection := client.Database("main").Collection("user")

    // create filter that's used to find the user
    filter := bson.D{{"uid", "Jacob"}}
    update := bson.D{
        {"$set", bson.D{
            {"sid", "this was updated"},
        }},
    }
    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }
    // if it didn't find any results when trying to update then add the user
    if updateResult.MatchedCount == 0 {
        fmt.Printf("didn't match any documents creating new one\n")
        //jacob := Users{"Jacob", "random fake id for a user"}
        type MongoFields struct {
          FieldStr string `json:"Field Str"`
          FieldInt string `json:"Field Int"`
        }

        jacob := MongoFields{
          FieldStr: "Jacob",
          FieldInt: "Random fake id for user",
        }

        insertResult, err := collection.InsertOne(context.TODO(), jacob)
        if err != nil {
          fmt.Printf("There was an err\n")
            log.Fatal(err)
        }
        fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
    } else {
        fmt.Printf("updated successfully\n")
    }

    err = client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}
