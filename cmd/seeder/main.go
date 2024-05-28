package main

import (
	"skripsi-be/cmd/seeder/seed"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)

	// _ = mongo

	seed.SeedUsers(mongo)
	seed.SeedRolesAndPermissions(mongo)
	seed.SeedOAuthScopes(mongo)
	seed.SeedOAuthClients(mongo)
}
