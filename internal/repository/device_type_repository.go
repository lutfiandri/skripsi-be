package repository

import (
	"context"

	"skripsi-be/internal/model/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceTypeRepository interface {
	GetDeviceTypes(ctx context.Context) ([]db.DeviceType, error)
	GetDeviceTypeById(ctx context.Context, id primitive.ObjectID) (db.DeviceType, error)
	UpsertDeviceType(ctx context.Context, id primitive.ObjectID, deviceType db.DeviceType) error
	DeleteDeviceType(ctx context.Context, id primitive.ObjectID) error
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

func (repository *deviceTypeRepository) GetDeviceTypes(ctx context.Context) ([]db.DeviceType, error) {
	var deviceTypes []db.DeviceType

	filter := bson.M{}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return deviceTypes, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var deviceType db.DeviceType
		if err := cursor.Decode(&deviceType); err != nil {
			return deviceTypes, err
		}

		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes, nil
}

func (repository *deviceTypeRepository) GetDeviceTypeById(ctx context.Context, id primitive.ObjectID) (db.DeviceType, error) {
	var deviceType db.DeviceType

	filter := bson.M{"_id": id}

	if err := repository.collection.FindOne(ctx, filter).Decode(&deviceType); err != nil {
		return deviceType, err
	}

	return deviceType, nil
}

func (repository *deviceTypeRepository) UpsertDeviceType(ctx context.Context, id primitive.ObjectID, deviceType db.DeviceType) error {
	filter := bson.M{"_id": id}
	update := deviceType
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *deviceTypeRepository) DeleteDeviceType(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
