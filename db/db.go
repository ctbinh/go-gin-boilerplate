package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc){
            
    // CancelFunc to cancel to context
    defer cancel()
     
    // client provides a method to close
    // a mongoDB connection.
    defer func(){
     
        // client.Disconnect method also has deadline.
        // returns error if any,
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}

func ping(client *mongo.Client, ctx context.Context) error{
 
    // mongo.Client has Ping to ping mongoDB, deadline of
    // the Ping method will be determined by cxt
    // Ping method return error if any occurred, then
    // the error can be handled.
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	mongoClient = client
	defer close(client, ctx, cancel)
	ping(client, ctx)
}

func GetDB() *mongo.Client {
	return mongoClient
}
