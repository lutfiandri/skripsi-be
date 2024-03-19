package repository

import (
	"context"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OAuthScopeRepository interface {
	GetOAuthScopes(ctx context.Context) ([]domain.OAuthScope, error)
	GetOAuthScopeById(ctx context.Context, id string) (domain.OAuthScope, error)
	UpsertOAuthScope(ctx context.Context, id string, oauthScope domain.OAuthScope) error
	DeleteOAuthScope(ctx context.Context, id string) error
}

type oauthScopeRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewOAuthScopeRepository(database *mongo.Database, collectionName string) OAuthScopeRepository {
	return &oauthScopeRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *oauthScopeRepository) GetOAuthScopes(ctx context.Context) ([]domain.OAuthScope, error) {
	var oauthScopes []domain.OAuthScope

	filter := bson.M{}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return oauthScopes, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var oauthScope domain.OAuthScope
		if err := cursor.Decode(&oauthScope); err != nil {
			return oauthScopes, err
		}

		oauthScopes = append(oauthScopes, oauthScope)
	}

	return oauthScopes, nil
}

func (repository *oauthScopeRepository) GetOAuthScopeById(ctx context.Context, id string) (domain.OAuthScope, error) {
	var oauthScope domain.OAuthScope

	filter := bson.M{"_id": id}

	if err := repository.collection.FindOne(ctx, filter).Decode(&oauthScope); err != nil {
		return oauthScope, err
	}

	return oauthScope, nil
}

func (repository *oauthScopeRepository) UpsertOAuthScope(ctx context.Context, id string, oauthScope domain.OAuthScope) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": oauthScope}
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *oauthScopeRepository) DeleteOAuthScope(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := repository.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
