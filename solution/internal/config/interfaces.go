package config

type HTTPConfig interface {
	Address() string
}

type ADMINConfig interface {
	ADMIN_EMAIL() string
	ADMIN_FULLNAME() string
	ADMIN_PASSWORD() string
}

type DATABASEConfig interface {
	DB_HOST() string
	DB_PORT() string
	DB_NAME() string
	DB_USER() string
	DB_PASSWORD() string
	URI() string
	MigrationsPath() string
}

type RANDOMSECRETSConfig interface {
	RANDOM_SECRET() string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
