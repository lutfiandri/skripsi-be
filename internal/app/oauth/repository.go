package oauth

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	// oauth_auth_code
	GetAuthCodeByCode(ctx context.Context, code string) (domain.OAuthAuthCode, error)
	InsertAuthCode(ctx context.Context, authCode domain.OAuthAuthCode) error
	DeleteAuthCodeByCode(ctx context.Context, code string) error

	// oauth_client
	GetOAuthClientById(ctx context.Context, id uuid.UUID) (domain.OAuthClient, error)

	// user
	GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error)
}

type repository struct {
	database                *mongo.Database
	oauthAuthCodeCollection *mongo.Collection
	oauthClientCollection   *mongo.Collection
	userCollection          *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	oauthAuthCodeCollection := database.Collection(domain.OAuthAuthCodeCollection)
	oauthAuthCodeCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{"auth_code": 1},
		},
		{
			Keys:    bson.M{"created_at": 1},
			Options: options.Index().SetExpireAfterSeconds(30),
		},
	})

	return &repository{
		database:                database,
		oauthAuthCodeCollection: oauthAuthCodeCollection,
		oauthClientCollection:   database.Collection(domain.OAuthClientCollection),
		userCollection:          database.Collection(domain.UserCollection),
	}
}

// oauth_auth_code
func (repository repository) GetAuthCodeByCode(ctx context.Context, code string) (domain.OAuthAuthCode, error) {
	var authCode domain.OAuthAuthCode

	filter := bson.M{"auth_code": code}

	if err := repository.oauthAuthCodeCollection.FindOne(ctx, filter).Decode(&authCode); err != nil {
		return authCode, err
	}

	return authCode, nil
}

func (repository repository) InsertAuthCode(ctx context.Context, authCode domain.OAuthAuthCode) error {
	if _, err := repository.oauthAuthCodeCollection.InsertOne(ctx, authCode); err != nil {
		return err
	}
	return nil
}

func (repository repository) DeleteAuthCodeByCode(ctx context.Context, code string) error {
	filter := bson.M{"auth_code": code}

	if _, err := repository.oauthAuthCodeCollection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

// oauth_client
func (repository repository) GetOAuthClientById(ctx context.Context, id uuid.UUID) (domain.OAuthClient, error) {
	var oauthClient domain.OAuthClient

	filter := bson.M{"_id": id}

	if err := repository.oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return oauthClient, err
	}

	return oauthClient, nil
}

// user
func (repository repository) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User

	filter := bson.M{"_id": id}

	if err := repository.userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}
