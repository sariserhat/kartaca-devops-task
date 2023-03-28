package repository

import (
	"context"
	"github.com/serhhatsari/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
)

var (
	ConnectionURI  = os.Getenv("CONNECTION_URI")
	DbName         = os.Getenv("DB_NAME")
	CollectionName = os.Getenv("COLLECTION_NAME")
)

type Repository struct {
	client *mongo.Client
}

func New() *Repository {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConnectionURI))
	if err != nil {
		panic(err)
	}
	return &Repository{
		client: client,
	}
}


func (r *Repository) GetRandomCountry() (model.Ulkeler, error) {
	collection := r.client.Database(DbName).Collection(CollectionName)

	count := getDocCount(collection)

	skip := rand.Int63n(count)

	var country model.Ulkeler

	err := collection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSkip(skip)).Decode(&country)
	if err != nil {
		return model.Ulkeler{}, err
	}
	return country, nil
}

func getDocCount(collection *mongo.Collection) int64 {
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	return count
}
