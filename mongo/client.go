package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitializeDatabase initializes MongoDB database client
func InitializeDatabase(uri string, db string) (database *mongo.Database, disconnect func() error, err error) {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	client, err := mongo.Connect(timeoutContext, options.Client().ApplyURI(uri))

	if err != nil {
		client.Disconnect(timeoutContext)
		cancel()
		return nil, nil, err
	}

	if err := client.Ping(timeoutContext, readpref.Primary()); err != nil {
		client.Disconnect(timeoutContext)
		cancel()
		return nil, nil, err
	}

	database = client.Database(db)

	return database, func() error {
		err := client.Disconnect(timeoutContext)
		cancel()
		return err
	}, nil
}
