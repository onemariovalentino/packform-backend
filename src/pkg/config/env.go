package config

import (
	"context"
	"log"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type (
	AppConfig struct {
		Name        string `env:"APP_NAME"`
		Environment string `env:"APP_ENV"`
		Port        int    `env:"APP_PORT"`
		Platform    string `env:"APP_PLATFORM"`
		DbConfig    *DbConfig
	}

	DbConfig struct {
		DbHost              string `env:"DB_HOST"`
		DbPort              int    `env:"DB_PORT"`
		DbName              string `env:"DB_NAME"`
		DbUsername          string `env:"DB_USERNAME"`
		DbPassword          string `env:"DB_PASSWORD"`
		DbSslMode           string `env:"DB_SSL_MODE"`
		DbMaxOpenConnection int    `env:"DB_MAX_OPEN_CONNECTION"`
		DbMaxIdleConnection int    `env:"DB_MAX_IDLE_CONNECTION"`
	}
)

var Env AppConfig

func LoadEnvironment() {
	ctx := context.Background()
	if err := envconfig.Process(ctx, &Env); err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(2)
	}
}
