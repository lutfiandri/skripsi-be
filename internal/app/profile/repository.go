package profile

import (
	"context"

	"skripsi-be/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
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

func (repository *repository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	filter := bson.M{"email": email}

	if err := repository.userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (repository *repository) UpdateUser(ctx context.Context, user domain.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}

	_, err := repository.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
