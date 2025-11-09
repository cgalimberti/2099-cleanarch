package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver      string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	WebServerPort string
	GRPCServerPort string
	GraphQLServerPort string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path + "/.env"); err != nil {
		return Config{}, err
	}

	return Config{
		DBDriver:      os.Getenv("DB_DRIVER"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
		WebServerPort: os.Getenv("WEB_SERVER_PORT"),
		GRPCServerPort: os.Getenv("GRPC_SERVER_PORT"),
		GraphQLServerPort: os.Getenv("GRAPHQL_SERVER_PORT"),
	}, nil
}