package controller

import (
	"context"
	"fmt"
	"time"

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

