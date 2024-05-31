package role

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetRoles(ctx context.Context) ([]domain.Role, error)
	GetRoleById(ctx context.Context, id uuid.UUID) (domain.Role, error)
	CreateRole(ctx context.Context, oauthClient domain.Role) error
	UpdateRole(ctx context.Context, oauthClient domain.Role) error
	DeleteRole(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	database              *mongo.Database
	oauthClientCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:              database,
		oauthClientCollection: database.Collection(domain.RoleCollection),
	}
}

func (repository repository) GetRoles(ctx context.Context) ([]domain.Role, error) {
	var oauthClients []domain.Role

	filter := bson.M{}

	cursor, err := repository.oauthClientCollection.Find(ctx, filter)
	if err != nil {
		return oauthClients, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var oauthClient domain.Role
		if err := cursor.Decode(&oauthClient); err != nil {
			return oauthClients, err
		}

		oauthClients = append(oauthClients, oauthClient)
	}

	return oauthClients, nil
}

func (repository repository) GetRoleById(ctx context.Context, id uuid.UUID) (domain.Role, error) {
	var oauthClient domain.Role

	filter := bson.M{"_id": id}

	if err := repository.oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return oauthClient, err
	}

	return oauthClient, nil
}

func (repository repository) CreateRole(ctx context.Context, oauthClient domain.Role) error {
	_, err := repository.oauthClientCollection.InsertOne(ctx, oauthClient)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateRole(ctx context.Context, oauthClient domain.Role) error {
	filter := bson.M{"_id": oauthClient.Id}
	update := bson.M{"$set": oauthClient}

	_, err := repository.oauthClientCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeleteRole(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := repository.oauthClientCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
