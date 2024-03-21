package oauth_scope

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetOAuthScopes(ctx context.Context) ([]domain.OAuthScope, error)
	GetOAuthScopeById(ctx context.Context, id uuid.UUID) (domain.OAuthScope, error)
	CreateOAuthScope(ctx context.Context, oauthScope domain.OAuthScope) error
	UpdateOAuthScope(ctx context.Context, oauthScope domain.OAuthScope) error
	DeleteOAuthScope(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	database             *mongo.Database
	oauthScopeCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:             database,
		oauthScopeCollection: database.Collection(domain.OAuthScopeCollection),
	}
}

func (repository repository) GetOAuthScopes(ctx context.Context) ([]domain.OAuthScope, error) {
	var oauthScopes []domain.OAuthScope

	filter := bson.M{}

	cursor, err := repository.oauthScopeCollection.Find(ctx, filter)
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

func (repository repository) GetOAuthScopeById(ctx context.Context, id uuid.UUID) (domain.OAuthScope, error) {
	var oauthScope domain.OAuthScope

	filter := bson.M{"_id": id}

	if err := repository.oauthScopeCollection.FindOne(ctx, filter).Decode(&oauthScope); err != nil {
		return oauthScope, err
	}

	return oauthScope, nil
}

func (repository repository) CreateOAuthScope(ctx context.Context, oauthScope domain.OAuthScope) error {
	_, err := repository.oauthScopeCollection.InsertOne(ctx, oauthScope)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) UpdateOAuthScope(ctx context.Context, oauthScope domain.OAuthScope) error {
	filter := bson.M{"_id": oauthScope.Id}

	_, err := repository.oauthScopeCollection.UpdateOne(ctx, filter, oauthScope)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) DeleteOAuthScope(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := repository.oauthScopeCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
