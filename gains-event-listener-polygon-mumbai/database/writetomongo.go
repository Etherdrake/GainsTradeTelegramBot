package database

import (
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func WriteTradeOrOrderToMongo(collectionName string, data interface{}) error {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Set up MongoDB collection
	collection := client.Database(gnsconstants.DbNameMumbai).Collection(collectionName)

	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert document into the collection
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	fmt.Println("Document inserted successfully")
	return nil
}

func WriteTradeOrOrderToMongoMarket(collectionName string, marketExec transformers.MarketExecutedTransform, open transformers.MarketOpenCanceledTransform, close transformers.MarketCloseCanceledTransform) error {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Println(err)
		}
	}()
	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check which type is non-nil and insert into the respective collection
	if marketExec != (transformers.MarketExecutedTransform{}) {
		// Set up MongoDB collection
		collection := client.Database(gnsconstants.DbNameMumbai).Collection(collectionName)

		// Insert document into the collection
		_, err = collection.InsertOne(ctx, marketExec)
		if err != nil {
			return err
		}
	} else if open != (transformers.MarketOpenCanceledTransform{}) {
		// Set up MongoDB collection
		collection := client.Database(gnsconstants.DbNameMumbai).Collection(collectionName)

		_, err = collection.InsertOne(ctx, open)
		if err != nil {
			return err
		}
	} else if close != (transformers.MarketCloseCanceledTransform{}) {
		// Set up MongoDB collection
		collection := client.Database(gnsconstants.DbNameMumbai).Collection(collectionName)

		_, err = collection.InsertOne(ctx, close)
		if err != nil {
			return err
		}
	} else {
		return errors.New("all input types are nil")
	}

	fmt.Println("Document inserted successfully")
	return nil
}
