package env

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type DatabaseEnvConfig struct {
	DB_HOST     string `env:"DB_HOST,required"`
	DB_PORT     string `env:"DB_PORT,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
}

type databaseConfig struct {
	raw DatabaseEnvConfig
}

func NewDatabaseConfig() (*databaseConfig, error) {
	var raw DatabaseEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &databaseConfig{raw: raw}, nil
}

func (cfg *databaseConfig) DB_HOST() string {
	return cfg.raw.DB_HOST
}

func (cfg *databaseConfig) DB_PORT() string {
	return cfg.raw.DB_PORT
}

func (cfg *databaseConfig) DB_NAME() string {
	return cfg.raw.DB_NAME
}

func (cfg *databaseConfig) DB_USER() string {
	return cfg.raw.DB_USER
}

func (cfg *databaseConfig) DB_PASSWORD() string {
	return cfg.raw.DB_PASSWORD
}

func (cfg *databaseConfig) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.raw.DB_USER,
		cfg.raw.DB_PASSWORD,
		cfg.raw.DB_HOST,
		cfg.raw.DB_PORT,
		cfg.raw.DB_NAME,
	)
}

func (cfg *databaseConfig) MigrationsPath() string {
	return "/opt/migrations"
}
