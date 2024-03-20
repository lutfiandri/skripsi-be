package device_type

import (
	"context"
	"time"

	"skripsi-be/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetDeviceTypes(ctx context.Context) ([]domain.DeviceType, error)
	GetDeviceTypeById(ctx context.Context, id string) (domain.DeviceType, error)
	CreateDeviceType(ctx context.Context, deviceType domain.DeviceType) error
	UpdateDeviceType(ctx context.Context, deviceType domain.DeviceType) error
	DeleteDeviceType(ctx context.Context, id string) error
}

type repository struct {
	database             *mongo.Database
	deviceTypeCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:             database,
		deviceTypeCollection: database.Collection(domain.DeviceTypeCollection),
	}
}

func (repository repository) GetDeviceTypes(ctx context.Context) ([]domain.DeviceType, error) {
	var deviceTypes []domain.DeviceType

	filter := bson.M{"deleted_at": nil}

	cursor, err := repository.deviceTypeCollection.Find(ctx, filter)
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

func (repository repository) GetDeviceTypeById(ctx context.Context, id string) (domain.DeviceType, error) {
	var deviceType domain.DeviceType

	filter := bson.M{"_id": id, "deleted_at": nil}

	if err := repository.deviceTypeCollection.FindOne(ctx, filter).Decode(&deviceType); err != nil {
		return deviceType, err
	}

	return deviceType, nil
}

func (repository repository) CreateDeviceType(ctx context.Context, deviceType domain.DeviceType) error {
	_, err := repository.deviceTypeCollection.InsertOne(ctx, deviceType)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateDeviceType(ctx context.Context, deviceType domain.DeviceType) error {
	filter := bson.M{"_id": deviceType.Id, "deleted_at": nil}

	update := bson.M{"$set": deviceType}

	_, err := repository.deviceTypeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeleteDeviceType(ctx context.Context, id string) error {
	filter := bson.M{"_id": id, "deleted_at": nil}
	update := bson.M{"$set": bson.M{
		"deleted_at": time.Now(),
	}}

	_, err := repository.deviceTypeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
