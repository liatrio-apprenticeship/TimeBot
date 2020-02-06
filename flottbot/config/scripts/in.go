package main

import (
    "fmt"
    "log"
    "time"
    "os"

	"golang.org/x/net/context"

    "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Time struct {
    Timestamp time.Time
    In bool
}

func main() {
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
    collection := mongoclient.Database("timesheets").Collection(os.Args[1])

    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"timestamp", -1}})
    findOptions.SetLimit(1)

    // Passing bson.D{{}} as the filter matches all documents in the collection
    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
    if err != nil {
        log.Fatal(err)
    }

    if cur.Next(context.TODO()) {
        // create a value into which the single document can be decoded
        var elem Time

        err = cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }

        if err := cur.Err(); err != nil {
            log.Fatal(err)
        }

        if (elem.In) {
            log.Fatal("Please use 'out' before using 'in'.")
        }
    }

    time_in := Time{time.Now(), true}
    insertResult, err := collection.InsertOne(context.TODO(), time_in)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    fmt.Println("Clocked In")

}
