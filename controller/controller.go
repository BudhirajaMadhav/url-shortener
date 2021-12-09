package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/budhirajamadhav/url-shortener/model"
	"github.com/budhirajamadhav/url-shortener/rand"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const connectionString = "mongodb+srv://<project>:<password>@cluster0.1dhwv.mongodb.net/url-shortener-db?retryWrites=true&w=majority"

const dbName = "url-shortener-db"
const colName = "pathsToUrls"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected Successfully")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))
	fmt.Println("Collection instance made successfully")

}

func shortenUrl(urlPath model.ShortenedUrl) string {
	insertResult, err := collection.InsertOne(context.Background(), urlPath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Url shortened with id", insertResult.InsertedID, "with path of", urlPath.Path)

	return urlPath.Path

}

func Redirector(fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/shorten" {
			fallback.ServeHTTP(w, r)
			return
		}

		var urlPath model.ShortenedUrl
		fmt.Println(path)
		err := collection.FindOne(context.Background(), bson.D{{"path", path}}).Decode(&urlPath)
		fmt.Println(urlPath, "urlpath")
		if err != nil {
			w.Write([]byte("Path does not exist"))
			return
		} else {
			http.Redirect(w, r, urlPath.URL, http.StatusFound)
			return
		}
	}
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var urlPath model.ShortenedUrl
	_ = json.NewDecoder(r.Body).Decode(&urlPath)
	
	if urlPath.Path == "" {
		urlPath.Path = rand.String(6)
	} else if pathStillExists(urlPath.Path) {
		alreadyExistsError := "<h1><span id=\"errorSpan\" style=\"color:red;\">Path still exists, choose another or use random by not mentioning any path</span></h1>"
		w.Write([]byte(alreadyExistsError))
		return
	}

	urlPath.Path = "/" + urlPath.Path

	shortenUrl(urlPath)

	w.Write([]byte(fmt.Sprintf("<h1>Your url is shortened to %s</h1>", urlPath.Path)))
}

func pathStillExists(path string) bool {
	
	var urlPath model.ShortenedUrl
	err := collection.FindOne(context.Background(), bson.D{{"path", path}}).Decode(&urlPath)
	if err != nil {
		return true
	}

	return false

}