package repository

import (
	"context"
	"hexago/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDDRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBRepository(database string, collection string) *MongoDDRepository {
	// TODO: URI should be extracted from env.
	opts := options.Client().ApplyURI("mongodb://username:password@localhost:27017/")

	// TODO: Create custom context ¿?
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		// panic(err)
	}

	// TODO: Defer was fucking everything up, how do we kill this thing then ¿?
	/*
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			} else {
				fmt.Println("Client is clossing")
			}
		}()
	*/

	coll := client.Database(database).Collection(collection)
	return &MongoDDRepository{client: client, collection: coll}
}

func (m *MongoDDRepository) GetFeatureFlag(name string) (*domain.FeatureFlagT, error) {
	var result domain.FeatureFlagT

	filter := bson.D{{Key: "name", Value: name}}
	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}
