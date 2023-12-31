package repository

import (
	"context"

	"skripsi-be/internal/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpsertUser(ctx context.Context, id string, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	filter := bson.M{}

	cursor, err := repository.collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *userRepository) GetUserById(ctx context.Context, id string) (domain.User, error) {
	var user domain.User

	filter := bson.M{"_id": id}

	if err := repository.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	filter := bson.M{"email": email}

	if err := repository.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepository) UpsertUser(ctx context.Context, id string, user domain.User) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	opts := options.Update().SetUpsert(true)

	_, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) DeleteUser(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := repository.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
