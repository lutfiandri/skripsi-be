package auth

import (
	"context"

	"skripsi-be/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

type repository struct {
	database       *mongo.Database
	userCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:       database,
		userCollection: database.Collection(domain.UserCollection),
	}
}

func (repository repository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	filter := bson.M{"email": email}

	if err := repository.userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (repository repository) CreateUser(ctx context.Context, user domain.User) error {
	_, err := repository.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
