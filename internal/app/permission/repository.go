package permission

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetPermissions(ctx context.Context) ([]domain.Permission, error)
	GetPermissionById(ctx context.Context, id uuid.UUID) (domain.Permission, error)
	CreatePermission(ctx context.Context, oauthClient domain.Permission) error
	UpdatePermission(ctx context.Context, oauthClient domain.Permission) error
	DeletePermission(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	database              *mongo.Database
	oauthClientCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:              database,
		oauthClientCollection: database.Collection(domain.PermissionCollection),
	}
}

func (repository repository) GetPermissions(ctx context.Context) ([]domain.Permission, error) {
	var oauthClients []domain.Permission

	filter := bson.M{}

	cursor, err := repository.oauthClientCollection.Find(ctx, filter)
	if err != nil {
		return oauthClients, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var oauthClient domain.Permission
		if err := cursor.Decode(&oauthClient); err != nil {
			return oauthClients, err
		}

		oauthClients = append(oauthClients, oauthClient)
	}

	return oauthClients, nil
}

func (repository repository) GetPermissionById(ctx context.Context, id uuid.UUID) (domain.Permission, error) {
	var oauthClient domain.Permission

	filter := bson.M{"_id": id}

	if err := repository.oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return oauthClient, err
	}

	return oauthClient, nil
}

func (repository repository) CreatePermission(ctx context.Context, oauthClient domain.Permission) error {
	_, err := repository.oauthClientCollection.InsertOne(ctx, oauthClient)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdatePermission(ctx context.Context, oauthClient domain.Permission) error {
	filter := bson.M{"_id": oauthClient.Id}
	update := bson.M{"$set": oauthClient}

	_, err := repository.oauthClientCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeletePermission(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := repository.oauthClientCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
