package main

import (
	"skripsi-be/cmd/seeder/seed"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)
	// seedOauthScopes(mongo)

	seed.SeedUsers(mongo)
	seed.SeedRolesAndPermissions(mongo)
}

// func seedOauthScopes(mongo *mongo.Database) {
// 	data := []interface{}{
// 		domain.OAuthScope{
// 			Id:          uuid.New(),
// 			Section:     "Device State",
// 			Code:        "read:device_state",
// 			Description: "Read device states",
// 		},
// 		domain.OAuthScope{
// 			Id:          uuid.New(),
// 			Section:     "Device State",
// 			Code:        "write:device_state",
// 			Description: "Write device states",
// 		},
// 	}

// 	mongo.Collection("oauth_scopes").InsertMany(context.Background(), data)
// }
