package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hadarshavit/local-opt/runner/results"
)

// MongoAdapter Adapeter for MongoDB
type MongoAdapter struct {
	client *mongo.Client
	db *mongo.Database
	resultsCollection *mongo.Collection
	problemsCollection *mongo.Collection
}

// NewMongoAdapter constructor
func NewMongoAdapter() MongoAdapter {
	return MongoAdapter{}
}

// Connect to db
func (adapter *MongoAdapter) Connect(ctx context.Context) error {
	err := adapter.dbConnect(ctx)
	if err != nil { return err }

	adapter.openLocalOptDB()
	adapter.openResultsCollection()
	adapter.openProblemsCollection()

	return nil
}

func (adapter *MongoAdapter) dbConnect(ctx context.Context) error {
	uri := "mongodb+srv://cluster0.tpdut.mongodb.net/local-optk;?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=runner/database/mongo.pem"
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil { return err }

	adapter.client = client

	return err
}

func (adapter *MongoAdapter) openLocalOptDB() {
	db := adapter.client.Database("local_opt")

	adapter.db = db
}

func (adapter *MongoAdapter) openResultsCollection() {
	collection := adapter.db.Collection("results")

	adapter.resultsCollection = collection
}

func (adapter *MongoAdapter) openProblemsCollection() {
	collection := adapter.db.Collection("problems")

	adapter.problemsCollection = collection
}

// SaveRunResults save one run results
func (adapter *MongoAdapter) SaveRunResults(ctx context.Context, results results.RunResults) (*mongo.InsertOneResult, error) {
	return adapter.resultsCollection.InsertOne(ctx, results)
}
