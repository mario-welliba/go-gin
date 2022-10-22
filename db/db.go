package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnection struct {
	client *mongo.Client
}

func NewConnection(host string) (conn *DBConnection) {
	log.Printf("Hello")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	return &DBConnection{client}
}

// Use handles connect to a certain collection
func (conn *DBConnection) Use(dbName, tableName string) (collection *mongo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.client.Database(dbName).Collection(tableName)
}

// Close handles closing a database connection
// func (conn *DBConnection) Close() {
// 	// This closes the connection
// 	conn.client.Disconnect()

// }
