package repository

import (
	"context"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OAuthAuthCodeRepository interface {
	GetAuthCodeByCode(ctx context.Context, code string) (domain.OAuthAuthCode, error)
	InsertAuthCode(ctx context.Context, authCode domain.OAuthAuthCode) error
	DeleteAuthCodeByCode(ctx context.Context, code string) error
}

type oauthAuthCodeRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewOAuthAuthCodeRepository(database *mongo.Database, collectionName string) OAuthAuthCodeRepository {
	collection := database.Collection(collectionName)
	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{"auth_code": 1},
		},
		{
			Keys:    bson.M{"created_at": 1},
			Options: options.Index().SetExpireAfterSeconds(30),
		},
	})

	return &oauthAuthCodeRepository{
		database:   database,
		collection: collection,
	}
}

func (repository oauthAuthCodeRepository) GetAuthCodeByCode(ctx context.Context, code string) (domain.OAuthAuthCode, error) {
	var authCode domain.OAuthAuthCode

	filter := bson.M{"auth_code": code}

	if err := repository.collection.FindOne(ctx, filter).Decode(&authCode); err != nil {
		return authCode, err
	}

	return authCode, nil
}

func (repository oauthAuthCodeRepository) InsertAuthCode(ctx context.Context, authCode domain.OAuthAuthCode) error {
	if _, err := repository.collection.InsertOne(ctx, authCode); err != nil {
		return err
	}
	return nil
}

func (repository oauthAuthCodeRepository) DeleteAuthCodeByCode(ctx context.Context, code string) error {
	filter := bson.M{"auth_code": code}

	if _, err := repository.collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
