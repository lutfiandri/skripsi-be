package seed

import (
	"context"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"
	"skripsi-be/internal/util/helper"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(mongo *mongo.Database) {
	roleAdminId := uuid.MustParse(constant.RoleAdminId)
	roleCustomerId := uuid.MustParse(constant.RoleCustomerId)

	adminPassword := "test1234"
	hashedAdminPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	helper.PanicIfErr(err)

	lutfiPassword := "test1234"
	hashedLutfiPassword, err := bcrypt.GenerateFromPassword([]byte(lutfiPassword), bcrypt.DefaultCost)
	helper.PanicIfErr(err)

	now := time.Now()

	users := []any{
		domain.User{
			Id:        uuid.New(),
			RoleId:    roleAdminId,
			Email:     "admin@mail.com",
			Name:      "Admin",
			Password:  string(hashedAdminPassword),
			CreatedAt: now,
			UpdatedAt: now,
		},
		domain.User{
			Id:        uuid.New(),
			RoleId:    roleCustomerId,
			Email:     "lutfiandri232@gmail.com",
			Name:      "Lutfi Andriyanto",
			Password:  string(hashedLutfiPassword),
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	mongo.Collection(domain.UserCollection).InsertMany(context.Background(), users)
}
