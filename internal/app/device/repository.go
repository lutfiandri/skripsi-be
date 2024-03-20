package device

import (
	"context"
	"time"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetDevices(ctx context.Context) ([]domain.Device, error)
	GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error)
	CreateDevice(ctx context.Context, device domain.Device) error
	UpdateDevice(ctx context.Context, device domain.Device) error
	DeleteDevice(ctx context.Context, id uuid.UUID) error
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

func (repository repository) GetDevices(ctx context.Context) ([]domain.Device, error) {
	var devices []domain.Device

	filter := bson.M{"deleted_at": nil}

	cursor, err := repository.deviceCollection.Find(ctx, filter)
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

func (repository repository) GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error) {
	var device domain.Device

	filter := bson.M{"_id": id, "deleted_at": nil}

	if err := repository.deviceCollection.FindOne(ctx, filter).Decode(&device); err != nil {
		return device, err
	}

	return device, nil
}

func (repository repository) CreateDevice(ctx context.Context, device domain.Device) error {
	_, err := repository.deviceCollection.InsertOne(ctx, device)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateDevice(ctx context.Context, device domain.Device) error {
	filter := bson.M{"_id": device.Id, "deleted_at": nil}

	update := bson.M{"$set": device}

	_, err := repository.deviceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeleteDevice(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id, "deleted_at": nil}
	update := bson.M{"$set": bson.M{
		"deleted_at": time.Now(),
	}}

	_, err := repository.deviceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
