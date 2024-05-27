package auth

import (
	"context"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

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
	database             *mongo.Database
	userCollection       *mongo.Collection
	roleCollection       *mongo.Collection
	permissionCollection *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repository{
		database:             database,
		userCollection:       database.Collection(domain.UserCollection),
		roleCollection:       database.Collection(domain.RoleCollection),
		permissionCollection: database.Collection(domain.PermissionCollection),
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
	role := domain.Role{}
	err := repository.roleCollection.FindOne(ctx, bson.M{"_id": roleId}).Decode(&role)
	helper.PanicIfErr(err)

	permissions := []domain.Permission{}
	cursor, err := repository.permissionCollection.Find(ctx, bson.M{"_id": bson.M{"$in": role.PermissionIds}})
	helper.PanicIfErr(err)

	err = cursor.All(ctx, &permissions)
	helper.PanicIfErr(err)

	return permissions, nil
}
