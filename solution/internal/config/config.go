package config

import (
	"solution/internal/config/env"
)

var appConfig *config

type config struct {
	Http         HTTPConfig
	Admin        ADMINConfig
	Database     DATABASEConfig
	RandomSecret RANDOMSECRETSConfig
	Logger       LoggerConfig
}

func Load() error {
	httpCfg, err := env.NewHTTPConfig()
	if err != nil {
		return err
	}
	databaseCfg, err := env.NewDatabaseConfig()
	if err != nil {
		return err
	}
	adminCfg, err := env.NewAdminConfig()
	if err != nil {
		return err
	}
	secretsCfg, err := env.NewSecretsConfig()
	if err != nil {
		return err
	}
	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}
	appConfig = &config{
		Http:         httpCfg,
		Admin:        adminCfg,
		Database:     databaseCfg,
		RandomSecret: secretsCfg,
		Logger:       loggerCfg,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
