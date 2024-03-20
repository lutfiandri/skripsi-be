package repository

import (
	"context"
	"time"

	"skripsi-be/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceRepository interface {
	GetDevices(ctx context.Context) ([]domain.Device, error)
	GetDeviceById(ctx context.Context, id string) (domain.Device, error)
	UpsertDevice(ctx context.Context, id string, device domain.Device) error
	UpdateUser(ctx context.Context, deviceId string, userId string) error
	DeleteDevice(ctx context.Context, id string) error
}

type deviceRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDeviceRepository(database *mongo.Database, collectionName string) DeviceRepository {
	return &deviceRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *deviceRepository) GetDevices(ctx context.Context) ([]domain.Device, error) {
	var devices []domain.Device

	filter := bson.M{"deleted_at": nil}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return devices, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var device domain.Device
		if err := cursor.Decode(&device); err != nil {
			return devices, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (repository *deviceRepository) GetDeviceById(ctx context.Context, id string) (domain.Device, error) {
	var device domain.Device

	filter := bson.M{"_id": id, "deleted_at": nil}

	if err := repository.collection.FindOne(ctx, filter).Decode(&device); err != nil {
		return device, err
	}

	return device, nil
}

func (repository *deviceRepository) UpsertDevice(ctx context.Context, id string, device domain.Device) error {
	filter := bson.M{"_id": id, "deleted_at": nil}

	update := bson.M{"$set": device}
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *deviceRepository) UpdateUser(ctx context.Context, deviceId string, userId string) error {
	filter := bson.M{"_id": deviceId, "deleted_at": nil}
	update := bson.M{"$set": bson.M{
		"user_id": userId,
	}}

	_, err := repository.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository *deviceRepository) DeleteDevice(ctx context.Context, id string) error {
	filter := bson.M{"_id": id, "deleted_at": nil}
	update := bson.M{"$set": bson.M{
		"deleted_at": time.Now(),
	}}

	_, err := repository.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
