package repository

import (
	"context"
	"log"
	"time"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceStateRepository interface {
	InsertDeviceState(ctx context.Context, state domain.DeviceState) error
	GetDeviceStates(ctx context.Context, from, to *time.Time, device_id *string) ([]domain.DeviceState, error)
}

type deviceStateRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDeviceStateRepository(database *mongo.Database, collectionName string) DeviceStateRepository {
	tsOptions := options.CreateCollection()
	metaField := "survey_id"
	granularity := "minutes"
	tsOptions.SetTimeSeriesOptions(&options.TimeSeriesOptions{
		TimeField:   "created_at",
		MetaField:   &metaField,
		Granularity: &granularity,
	})

	if err := database.CreateCollection(context.Background(), collectionName, tsOptions); err != nil {
		log.Printf("WARNING: error on creating %s timeseries collection: %s\n", collectionName, err.Error())
	} else {
		log.Printf("%s timeseries collection created\n", collectionName)
	}

	collection := database.Collection(collectionName)
	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.M{"_id": 1}},
		{Keys: bson.M{"device_id": 1}},
	})

	return &deviceStateRepository{
		database:   database,
		collection: collection,
	}
}

func (repository *deviceStateRepository) InsertDeviceState(ctx context.Context, state domain.DeviceState) error {
	if _, err := repository.collection.InsertOne(ctx, state); err != nil {
		return err
	}
	return nil
}

func (repository *deviceStateRepository) GetDeviceStates(ctx context.Context, from, to *time.Time, device_id *string) ([]domain.DeviceState, error) {
	var deviceStates []domain.DeviceState

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
