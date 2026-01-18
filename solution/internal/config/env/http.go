package env

import "github.com/caarlos0/env/v11"

type HTTPEnvConfig struct {
	Port string `env:"SERVER_PORT,required"`
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
	return "8080"
}
