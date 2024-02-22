package repository

import (
	"context"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OAuthClientRepository interface {
	GetOAuthClients(ctx context.Context) ([]domain.OAuthClient, error)
	GetOAuthClientById(ctx context.Context, id string) (domain.OAuthClient, error)
	UpsertOAuthClient(ctx context.Context, id string, oauthClient domain.OAuthClient) error
	DeleteOAuthClient(ctx context.Context, id string) error
}

type oauthClientRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewOAuthClientRepository(database *mongo.Database, collectionName string) OAuthClientRepository {
	return &oauthClientRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *oauthClientRepository) GetOAuthClients(ctx context.Context) ([]domain.OAuthClient, error) {
	var oauthClients []domain.OAuthClient

	filter := bson.M{}

	cursor, err := repository.collection.Find(ctx, filter)
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

func (repository *oauthClientRepository) GetOAuthClientById(ctx context.Context, id string) (domain.OAuthClient, error) {
	var oauthClient domain.OAuthClient

	filter := bson.M{"_id": id}

	if err := repository.collection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return oauthClient, err
	}

	return oauthClient, nil
}

func (repository *oauthClientRepository) UpsertOAuthClient(ctx context.Context, id string, oauthClient domain.OAuthClient) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": oauthClient}
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *oauthClientRepository) DeleteOAuthClient(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := repository.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
