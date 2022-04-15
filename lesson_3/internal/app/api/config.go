package api

import "lesson_3/storage"

// Config Конфиг для нашего приложения
type Config struct {
	BindAddr    string `toml:"bind_addr" env:"BIND_ADDR"`
	LoggerLevel string `toml:"logger_level" env:"LOGGER_LEVEL"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8001",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
