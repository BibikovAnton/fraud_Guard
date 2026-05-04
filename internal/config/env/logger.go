package env

type LoggerEnvConfig struct {
	Level  string
	AsJson bool
}

type loggerConfig struct {
	raw LoggerEnvConfig
}

func NewLoggerConfig() (*loggerConfig, error) {
	raw := LoggerEnvConfig{
		Level:  "info",
		AsJson: true,
	}

	return &loggerConfig{raw: raw}, nil
}

func (cfg *loggerConfig) Level() string { return cfg.raw.Level }

func (cfg *loggerConfig) AsJson() bool { return cfg.raw.AsJson }
