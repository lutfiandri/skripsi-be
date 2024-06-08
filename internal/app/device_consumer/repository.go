package device_consumer

import (
	"context"
	"log"
	"time"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertDeviceState(ctx context.Context, state domain.DeviceStateLog[any]) error
	GetDeviceStates(ctx context.Context, from, to *time.Time, device_id *string) ([]domain.DeviceStateLog[any], error)

	UpdateDeviceLastState(ctx context.Context, state domain.DeviceStateLog[any]) error

	GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error)
}

type repository struct {
	database                 *mongo.Database
	deviceStateLogCollection *mongo.Collection
	deviceCollection         *mongo.Collection
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
		log.Printf("WARNING: error on creating %s timeseries deviceStateLogCollection: %s\n", collectionName, err.Error())
	} else {
		log.Printf("%s timeseries deviceStateLogCollection created\n", collectionName)
	}

	deviceStateLogCollection := database.Collection(collectionName)
	deviceStateLogCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.M{"_id": 1}},
		{Keys: bson.M{"device_id": 1}},
		{Keys: bson.M{"created_at": 1}},
	})

	return &repository{
		database:                 database,
		deviceStateLogCollection: deviceStateLogCollection,
		deviceCollection:         database.Collection(domain.DeviceCollection),
	}
}

func (repository repository) InsertDeviceState(ctx context.Context, state domain.DeviceStateLog[any]) error {
	if _, err := repository.deviceStateLogCollection.InsertOne(ctx, state); err != nil {
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

	cursor, err := repository.deviceStateLogCollection.Find(ctx, filter)
	if err != nil {
		return deviceStates, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &deviceStates); err != nil {
		return deviceStates, err
	}

	return deviceStates, nil
}

func (repository repository) UpdateDeviceLastState(ctx context.Context, state domain.DeviceStateLog[any]) error {
	filter := bson.M{"_id": state.DeviceId}

	update := bson.D{{
		"$set", bson.D{
			{"last_state", state.State},
			{"updated_at", time.Now()},
		},
	}}

	_, err := repository.deviceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error) {
	var device domain.Device

	filter := bson.M{"_id": id, "deleted_at": nil}

	if err := repository.deviceCollection.FindOne(ctx, filter).Decode(&device); err != nil {
		return device, err
	}

	return device, nil
}
