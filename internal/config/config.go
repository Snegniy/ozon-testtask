package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DebugMode          string `env:"SERVER_LOG_MODE" env-description:"Debug mode logger" env-default:"yes"`
	HTTPServerHostPort string `env:"SERVER_HTTP_HOST_PORT" env-default:"localhost:8000"`
	GRPCServerHostPort string `env:"SERVER_GRPC_HOST_PORT" env-default:"localhost:9000"`
	StorageType        string `env:"STORAGE_TYPE" env-default:"memdb"`
	Postgres           Postgres
}

type Postgres struct {
	Username string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	HostPort string `env:"DB_HOST_PORT" env-default:"localhost:8888"`
}

var path = "config.env"

func NewConfig() Config {

	log.Println("\t\tRead application configuration...")
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		help, _ := cleanenv.GetDescription(&cfg, nil)
		log.Println(help)
		log.Fatalf("%s", err)
	}
	log.Println("\t\tGet configuration - OK!")

	return cfg
}
