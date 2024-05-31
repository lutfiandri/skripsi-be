package seed

import (
	"context"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedOAuthClients(mongo *mongo.Database) {
	googleHomeId := uuid.MustParse(constant.OAuthClientGoogleHomeId)

	scopeReadDeviceId := uuid.MustParse(constant.OAuthScopeReadDeviceId)
	scopeUpdateDeviceId := uuid.MustParse(constant.OAuthScopeUpdateDeviceId)

	now := time.Now()

	clients := []any{
		domain.OAuthClient{
			Id:     googleHomeId,
			Secret: "m7HyjOD1-jt57teSiRKwIEaZc86afXhgHc0oUr6YOrwAeTDmgT9-riOlG6K_91WH5rILnF_Eyvbd31qY2e-Rkg==",
			Name:   "Google",
			RedirectUris: []string{
				"https://oauth-redirect.googleusercontent.com/r/skripsi---lutfi-smarthom-44706",
				"https://oauth-redirect-sandbox.googleusercontent.com/r/skripsi---lutfi-smarthom-44706",
			},
			ScopeIds:  uuid.UUIDs{scopeReadDeviceId, scopeUpdateDeviceId},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	mongo.Collection(domain.OAuthClientCollection).InsertMany(context.Background(), clients)
}
