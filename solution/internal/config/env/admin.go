package env

import "github.com/caarlos0/env/v11"

type AdminEnvConfig struct {
	ADMIN_EMAIL    string `env:"ADMIN_EMAIL,required"`
	ADMIN_FULLNAME string `env:"ADMIN_FULLNAME,required"`
	ADMIN_PASSWORD string `env:"ADMIN_PASSWORD,required"`
}

type adminConfig struct {
	raw AdminEnvConfig
}

func NewAdminConfig() (*adminConfig, error) {
	var raw AdminEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &adminConfig{raw: raw}, nil
}

func (cfg *adminConfig) ADMIN_EMAIL() string {
	return cfg.raw.ADMIN_EMAIL
}
func (cfg *adminConfig) ADMIN_FULLNAME() string {
	return cfg.raw.ADMIN_FULLNAME
}
func (cfg *adminConfig) ADMIN_PASSWORD() string {
	return cfg.raw.ADMIN_PASSWORD
}
