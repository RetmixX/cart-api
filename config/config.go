package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type AppConf struct {
	ServerPort int    `env:"SERVER_PORT" env-required:"true"`
	ServerMode string `env:"SERVER_MODE" env-default:"local"`
	DBConf     DbConf
}

type DbConf struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Username string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
}

func MustLoad() *AppConf {
	configPath := ".deploy/local/.env"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file don't exists" + configPath)
	}

	var cfg AppConf

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config file: " + err.Error())
	}

	return &cfg
}
