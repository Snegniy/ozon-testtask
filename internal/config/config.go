package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	DebugMode          string `env:"SERVER_DEBUG_MODE" env-description:"Debug mode logger" env-default:"yes"`
	HTTPServerHostPort string `env:"SERVER_HTTP_HOST_PORT" env-default:"localhost:8000"`
	GRPCServerHostPort string `env:"SERVER_GRPC_HOST_PORT" env-default:"localhost:9000"`
	StorageType        string `env:"STORAGE_TYPE" env-default:""`
	Postgres           Postgres
}

type Postgres struct {
	Username string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
}

var path = ".env"

func NewConfig() Config {

	log.Println("\t\tRead application configuration...")
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		help, _ := cleanenv.GetDescription(&cfg, nil)
		log.Println(help)
		log.Fatalf("%s", err)
	}
	env := os.Getenv("STORAGE_TYPE")
	if env != "" {
		cfg.StorageType = env
	}
	log.Printf("\t\tSTORAGE_TYPE=%s\n", cfg.StorageType)
	log.Println("\t\tGet configuration - OK!")

	return cfg
}
