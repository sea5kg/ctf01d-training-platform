package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		DB   `yaml:"db"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT" env-default:"4102"`
		Host string `env-required:"true" yaml:"host" env:"HTTP_PORT" env-default:"localhost"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	DB struct {
		Driver     string `yaml:"driver"`
		DataSource string `env-required:"true" yaml:"data_source" env:"data_source"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) ParseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
