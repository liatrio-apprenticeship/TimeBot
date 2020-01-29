package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
    Uid string
    Sid string
}

func main() {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://database:27017")

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
    fmt.Println("user email: ", os.Args[1])
    collection := client.Database("main").Collection("user")

    // create filter that's used to find the user
    filter := bson.D{{"uid", os.Args[1]}}
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
        jacob := Users{os.Args[1], "random fake id for a user"}
        insertResult, err := collection.InsertOne(context.TODO(), jacob)
        if err != nil {
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
