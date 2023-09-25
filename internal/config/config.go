package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JWTSecretKey  = os.Getenv("JWT_SECRET_KEY")
	MONGO_URI     = os.Getenv("MONGO_URI")
	MONGO_DB_NAME = os.Getenv("MONGO_DB_NAME")
)
