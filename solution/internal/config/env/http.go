package env

import (
	"github.com/caarlos0/env/v11"
	"os"
)

type HTTPEnvConfig struct {
	Port string `env:"SERVER_PORT"`
}

type httpConfig struct {
	raw HTTPEnvConfig
}

func NewHTTPConfig() (*httpConfig, error) {
	var raw HTTPEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &httpConfig{raw: raw}, nil
}

func (cfg *httpConfig) Address() string {
	if cfg.raw.Port == "" {
		cfg.raw.Port = "8080"
	}
	return cfg.raw.Port
}
