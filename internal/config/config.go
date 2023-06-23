package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DebugMode      string `env:"SERVER_LOG_MODE" env-description:"Debug mode logger" env-default:"yes"`
	ServerHostPort string `env:"SERVER_HOST_PORT" env-default:"localhost:8000"`
	StorageType    string `env:"STORAGE_TYPE" env-default:"local"`
	TransportType  string `env:"TRANSPORT_TYPE" env-default:"http"`
	Postgres       Postgres
}

type Postgres struct {
	Username string `env:"PG_USERNAME" env-default:"postgres"`
	Password string `env:"PG_PASSWORD" env-default:"postgres"`
	Host     string `env:"PG_HOST" env-default:"localhost"`
	Port     string `env:"PG_PORT" env-default:"8888"`
	Database string `env:"PG_DATABASE" env-default:"links"`
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
