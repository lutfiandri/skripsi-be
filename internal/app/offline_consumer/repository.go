package offline_consumer

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UpdateDeviceOnline(ctx context.Context, deviceId uuid.UUID, isOnline bool) error
	GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error)
}

type repository struct {
	database         *mongo.Database
	deviceCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:         database,
		deviceCollection: database.Collection(domain.DeviceCollection),
	}
}

func (repository repository) UpdateDeviceOnline(ctx context.Context, deviceId uuid.UUID, isOnline bool) error {
	filter := bson.M{
		"_id":        deviceId,
		"deleted_at": nil,
	}

	update := bson.D{{
		"$set", bson.D{
			{"last_state.online", isOnline},
		},
	}}

	err := repository.deviceCollection.FindOneAndUpdate(ctx, filter, update).Err()
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
