package auth

import (
	"context"

	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error

	GetPermissionsByRoleId(ctx context.Context, roleId uuid.UUID) ([]domain.Permission, error)
}

type repository struct {
	database       *mongo.Database
	userCollection *mongo.Collection
	roleCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:       database,
		userCollection: database.Collection(domain.UserCollection),
		roleCollection: database.Collection(domain.RoleCollection),
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

func (repository repository) GetPermissionsByRoleId(ctx context.Context, roleId uuid.UUID) ([]domain.Permission, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"_id": roleId,
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         domain.PermissionCollection,
			"localField":   "permission_ids",
			"foreignField": "_id",
			"as":           "permissions",
		}}},
		{{Key: "$limit", Value: 1}},
	}

	// Execute the pipeline
	// note: aggregation always return limit
	cursor, err := repository.roleCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var roles []domain.Role
	if err = cursor.All(ctx, &roles); err != nil || len(roles) == 0 {
		return []domain.Permission{}, err
	}

	role := roles[0]

	return role.Permissions, nil
}
