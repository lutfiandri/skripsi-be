package device_consumer

import (
	"context"
	"log"
	"time"

	"skripsi-be/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertDeviceState(ctx context.Context, state domain.DeviceStateLog[any]) error
	GetDeviceStates(ctx context.Context, from, to *time.Time, device_id *string) ([]domain.DeviceStateLog[any], error)
}

type repository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	tsOptions := options.CreateCollection()
	metaField := "survey_id"
	granularity := "minutes"
	tsOptions.SetTimeSeriesOptions(&options.TimeSeriesOptions{
		TimeField:   "created_at",
		MetaField:   &metaField,
		Granularity: &granularity,
	})

	collectionName := domain.DeviceStateLogCollection

	if err := database.CreateCollection(context.Background(), collectionName, tsOptions); err != nil {
		log.Printf("WARNING: error on creating %s timeseries collection: %s\n", collectionName, err.Error())
	} else {
		log.Printf("%s timeseries collection created\n", collectionName)
	}

	collection := database.Collection(collectionName)
	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.M{"_id": 1}},
		{Keys: bson.M{"device_id": 1}},
		{Keys: bson.M{"created_at": 1}},
	})

	return &repository{
		database:   database,
		collection: collection,
	}
}

func (repository repository) InsertDeviceState(ctx context.Context, state domain.DeviceStateLog[any]) error {
	if _, err := repository.collection.InsertOne(ctx, state); err != nil {
		return err
	}
	return nil
}

func (repository repository) GetDeviceStates(ctx context.Context, from, to *time.Time, device_id *string) ([]domain.DeviceStateLog[any], error) {
	var deviceStates []domain.DeviceStateLog[any]

	filter := bson.M{}

	createdAtFilter := bson.M{}
	if from != nil {
		createdAtFilter["$gte"] = from
	}
	if to != nil {
		createdAtFilter["$lte"] = to
	}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return deviceStates, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &deviceStates); err != nil {
		return deviceStates, err
	}

	return deviceStates, nil
}
