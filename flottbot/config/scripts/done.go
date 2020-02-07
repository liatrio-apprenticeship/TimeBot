package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "time"
        //"math"

        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/sheets/v4"

        "go.mongodb.org/mongo-driver/bson"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)

// TODO put these structs in seperate files since they are used
// in multiple scripts
type Time struct {
    Timestamp time.Time
    In bool
    TimeSpent float64
    Day bool
}

type Users struct {
    Uid string
    Sid string
}

///////////Google Sheets Setup that is used in multiple files //////////

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "/tokens/token.json"
        tok, _ := tokenFromFile(tokFile)
        return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

///////////End Google Sheets Setup that is used in multiple files //////////

func main() {
    /////////////////  Google Sheets Setup  ///////////////////////
    b, err := ioutil.ReadFile("/tokens/credentials.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config)

    srv, err := sheets.New(client)
    _, err = sheets.New(client) // comment when using google sheets
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }
    /////////////////  End Google Sheets Setup  //////////////////////


    /////////////////  MongoDB Setup  ///////////////////////
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
    collectionSheet := mongoclient.Database("timesheets").Collection(os.Args[1])
    collectionUser := mongoclient.Database("main").Collection("users")
    /////////////////  End MongoDB Setup  ///////////////////////

    // Get the user's spreadsheet id from the database
    var cur_user Users
    filter := bson.D{{"uid", os.Args[1]}}

    err = collectionUser.FindOne(context.TODO(), filter).Decode(&cur_user)
    if err != nil {
        fmt.Println("We can't find your Google sheet in the database, use the set command to add it.")
        log.Fatal(err)
    }

    // Setup find options to onlt get the most recent entry in database
    //findOptions := options.Find()
    //findOptions.SetSort(bson.D{{"timestamp", -1}})
    //findOptions.SetLimit(1)

    filter2 := bson.D{{"day", true}}
    // Passing bson.D{{}} as the filter matches all documents in the collection
    cur, err := collectionSheet.Find(context.TODO(), filter2/*, findOptions*/)
    if err != nil {
        log.Fatal(err)
    }

    update := bson.M{
        "$set": bson.M{
        "day" : false,
        },
    }
    
    _, err = collectionSheet.UpdateMany(
        context.TODO(),
        filter2,
        update,
    )
    if err != nil {
        log.Fatal(err)
    }

    sum := 0.0

    for cur.Next(context.TODO()) {

        // create a value into which the single document can be decoded
        var elem Time
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }

        //result = append(results, &elem)
        sum += elem.TimeSpent
    }

    fmt.Println("Your Total: ", sum)

    spreadsheetId := cur_user.Sid
    writeRange := "Sheet1!A2:F2"

    var vr sheets.ValueRange

    myval := []interface{}{ nil, nil, nil, nil, nil, sum}
    vr.Values = append(vr.Values, myval)

    // Add new entry to end of Google Sheets document
    _, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet. %v", err)
    }

}
