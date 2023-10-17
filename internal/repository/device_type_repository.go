package repository

import (
	"context"
	"time"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceTypeRepository interface {
	GetDeviceTypes(ctx context.Context) ([]domain.DeviceType, error)
	GetDeviceTypeById(ctx context.Context, id string) (domain.DeviceType, error)
	UpsertDeviceType(ctx context.Context, id string, deviceType domain.DeviceType) error
	DeleteDeviceType(ctx context.Context, id string) error
}

type deviceTypeRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDeviceTypeRepository(database *mongo.Database, collectionName string) DeviceTypeRepository {
	return &deviceTypeRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *deviceTypeRepository) GetDeviceTypes(ctx context.Context) ([]domain.DeviceType, error) {
	var deviceTypes []domain.DeviceType

	filter := bson.M{"deleted_at": nil}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return deviceTypes, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var deviceType domain.DeviceType
		if err := cursor.Decode(&deviceType); err != nil {
			return deviceTypes, err
		}

		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes, nil
}

func (repository *deviceTypeRepository) GetDeviceTypeById(ctx context.Context, id string) (domain.DeviceType, error) {
	var deviceType domain.DeviceType

	filter := bson.M{"_id": id, "deleted_at": nil}

	if err := repository.collection.FindOne(ctx, filter).Decode(&deviceType); err != nil {
		return deviceType, err
	}

	return deviceType, nil
}

func (repository *deviceTypeRepository) UpsertDeviceType(ctx context.Context, id string, deviceType domain.DeviceType) error {
	filter := bson.M{"_id": id, "deleted_at": nil}

	update := bson.M{"$set": deviceType}
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *deviceTypeRepository) DeleteDeviceType(ctx context.Context, id string) error {
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
