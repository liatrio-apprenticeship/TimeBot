package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "time"
        "math"

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
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
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
    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"timestamp", -1}})
    findOptions.SetLimit(1)

    // Passing bson.D{{}} as the filter matches all documents in the collection
    cur, err := collectionSheet.Find(context.TODO(), bson.D{{}}, findOptions)
    if err != nil {
        log.Fatal(err)
    }

    // create a value into which the single document can be decoded
    var elem Time

    // ensure that the previous command used was in
    if cur.Next(context.TODO()) {
        // convert the results of the find command to a Time struct
        err = cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }

        if err := cur.Err(); err != nil {
            log.Fatal(err)
        }

        if (!elem.In) {
            log.Fatal("Please use 'in' before using 'out'.")
        }
    }
    // Get the current time
    cur_time := time.Now()
    // subtract the in and out times and convert to hours
    // then round the time to the nearest 2 decimals
    time_tot := math.Round(cur_time.Sub(elem.Timestamp).Hours()/0.25)*0.25

    time_out := Time{cur_time, false, time_tot}
    _, err = collectionSheet.InsertOne(context.TODO(), time_out)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("You Worked", time_tot, "Hours")

    spreadsheetId := cur_user.Sid
    writeRange := "Sheet1!B1:F1"

    var vr sheets.ValueRange

    myval := []interface{}{ cur_time.Format("01/02/2006"), nil,
         time_tot, elem.Timestamp.Format("01/02/2006 03:04:05 PM"),
         cur_time.Format("01/02/2006 03:04:05 PM")}
    vr.Values = append(vr.Values, myval)

    // Add new entry to end of Google Sheets document
    _, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet. %v", err)
    }

}
