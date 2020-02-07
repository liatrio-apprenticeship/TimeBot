package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "io"
    "io/ioutil"
    "regexp"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
    Uid string
    Sid string
}

func main() {

    content, err := ioutil.ReadFile("/go/mongodb/mongodb-root-password")

    fmt.Printf("File contents: %s", content)
    

    // // Set client options
    // clientOptions := options.Client().ApplyURI("mongodb://database:27017")

    // // Connect to MongoDB
    // client, err := mongo.Connect(context.TODO(), clientOptions)

    // if err != nil {
    //     log.Fatal(err)
    // }

    // // Check the connection
    // err = client.Ping(context.TODO(), nil)

    // if err != nil {
    //     log.Fatal(err)
    // }
    // // create regex that will pull out the sheetid from a url
    // re := regexp.MustCompile(`https://docs.google.com/spreadsheets/d/([a-zA-Z0-9-_]+)/edit[#&]gid=([0-9]+)`)
    // sheetid := re.FindSubmatch([]byte(os.Args[2]))
    // // if the user entered a bad url
    // if sheetid == nil {
    //     fmt.Println("Please enter valid Google Sheets url\n")
    //     os.Exit(1)
    // }
    // string_sheet_id := string(sheetid[1])
    // collection := client.Database("main").Collection("users")

    // // create filter that's used to find the user
    // filter := bson.D{{"uid", os.Args[1]}}
    // update := bson.D{
    //     {"$set", bson.D{
    //         {"sid", string_sheet_id},
    //     }},
    // }
    // // attempt to update their sheet
    // updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // // if it didn't find any results when trying to update then add the user
    // if updateResult.MatchedCount == 0 {
    //     user := Users{os.Args[1], string_sheet_id}
    //     insertResult, err := collection.InsertOne(context.TODO(), user)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     log.Print(insertResult)
    //     fmt.Println("Successfully added your sheet")
    // } else {
    //     fmt.Printf("Your sheet has been successfully updated\n")
    // }

    // err = client.Disconnect(context.TODO())

    // if err != nil {
    //     log.Fatal(err)
    // }

}
