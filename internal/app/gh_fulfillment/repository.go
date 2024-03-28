package gh_fulfillment

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetDevices(ctx context.Context, userId uuid.UUID, deviceIds *[]uuid.UUID) ([]domain.Device, error)
	GetDeviceById(ctx context.Context, id uuid.UUID) (domain.Device, error)
	UpdateDevice(ctx context.Context, device domain.Device) error

	UpdateDevicesLastState(ctx context.Context, deviceIds []uuid.UUID, state map[string]any) error

	GetDeviceTypes(ctx context.Context) ([]domain.DeviceType, error)
}

type repository struct {
	database             *mongo.Database
	deviceCollection     *mongo.Collection
	deviceTypeCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:             database,
		deviceCollection:     database.Collection(domain.DeviceCollection),
		deviceTypeCollection: database.Collection(domain.DeviceTypeCollection),
	}
}

func (repository repository) GetDevices(ctx context.Context, userId uuid.UUID, deviceIds *[]uuid.UUID) ([]domain.Device, error) {
	var devices []domain.Device

	filter := bson.M{
		"user_id":    userId,
		"deleted_at": nil,
	}

	if deviceIds != nil {
		filter["_id"] = bson.M{"$in": deviceIds}
	}

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

func (repository repository) UpdateDevice(ctx context.Context, device domain.Device) error {
	filter := bson.M{"_id": device.Id, "deleted_at": nil}

	update := bson.M{"$set": device}

	_, err := repository.deviceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateDevicesLastState(ctx context.Context, deviceIds []uuid.UUID, state map[string]any) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": deviceIds,
		},
		"deleted_at": nil,
	}

	updateState := map[string]any{}
	for key, value := range updateState {
		updateState["last_state."+key] = value
	}

	update := bson.M{"$set": updateState}

	_, err := repository.deviceCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
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
