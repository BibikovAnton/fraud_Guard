package env

import "github.com/caarlos0/env/v11"

type SecretsEnvConfig struct {
	RANDOM_SECRET string `env:"RANDOM_SECRET,required"`
}

type secretsConfig struct {
	raw SecretsEnvConfig
}

func NewSecretsConfig() (*secretsConfig, error) {
	var raw SecretsEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}
	return &secretsConfig{raw: raw}, nil
}

func (cfg *secretsConfig) RANDOM_SECRET() string {
	return cfg.raw.RANDOM_SECRET
}
