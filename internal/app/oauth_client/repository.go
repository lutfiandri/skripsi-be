package oauth_client

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetOAuthClients(ctx context.Context) ([]domain.OAuthClient, error)
	GetOAuthClientById(ctx context.Context, id uuid.UUID) (domain.OAuthClient, error)
	CreateOAuthClient(ctx context.Context, oauthClient domain.OAuthClient) error
	UpdateOAuthClient(ctx context.Context, oauthClient domain.OAuthClient) error
	DeleteOAuthClient(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	database              *mongo.Database
	oauthClientCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:              database,
		oauthClientCollection: database.Collection(domain.OAuthClientCollection),
	}
}

func (repository repository) GetOAuthClients(ctx context.Context) ([]domain.OAuthClient, error) {
	var oauthClients []domain.OAuthClient

	filter := bson.M{}

	cursor, err := repository.oauthClientCollection.Find(ctx, filter)
	if err != nil {
		return oauthClients, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var oauthClient domain.OAuthClient
		if err := cursor.Decode(&oauthClient); err != nil {
			return oauthClients, err
		}

		oauthClients = append(oauthClients, oauthClient)
	}

	return oauthClients, nil
}

func (repository repository) GetOAuthClientById(ctx context.Context, id uuid.UUID) (domain.OAuthClient, error) {
	var oauthClient domain.OAuthClient

	filter := bson.M{"_id": id}

	if err := repository.oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return oauthClient, err
	}

	return oauthClient, nil
}

func (repository repository) CreateOAuthClient(ctx context.Context, oauthClient domain.OAuthClient) error {
	_, err := repository.oauthClientCollection.InsertOne(ctx, oauthClient)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateOAuthClient(ctx context.Context, oauthClient domain.OAuthClient) error {
	filter := bson.M{"_id": oauthClient.Id}
	update := bson.M{"$set": oauthClient}

	_, err := repository.oauthClientCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeleteOAuthClient(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := repository.oauthClientCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
