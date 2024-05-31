package seed

import (
	"context"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedOAuthScopes(mongo *mongo.Database) {
	mainSeeds := getMainSeeds()

	seedOAuthScopes(mongo, mainSeeds)
}

func seedOAuthScopes(mongo *mongo.Database, mainSeeds []any) {
	permissionMap := make(map[uuid.UUID][]uuid.UUID)
	for _, s := range mainSeeds {
		seed := s.(PermissionSeed)

		for _, oauthScopeId := range seed.OAuthScopeIds {
			permissionMap[oauthScopeId] = append(permissionMap[oauthScopeId], seed.Permission.Id)
		}
	}

	scopeReadDeviceId := uuid.MustParse(constant.OAuthScopeReadDeviceId)
	scopeUpdateDeviceId := uuid.MustParse(constant.OAuthScopeUpdateDeviceId)
	now := time.Now()

	scopes := []any{
		domain.OAuthScope{
			Id:            scopeReadDeviceId,
			Description:   constant.OAuthScopeReadDeviceDescription,
			PermissionIds: permissionMap[scopeReadDeviceId],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		domain.OAuthScope{
			Id:            scopeUpdateDeviceId,
			Description:   constant.OAuthScopeUpdateDeviceDescription,
			PermissionIds: permissionMap[scopeUpdateDeviceId],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}

	mongo.Collection(domain.OAuthScopeCollection).InsertMany(context.Background(), scopes)
}
