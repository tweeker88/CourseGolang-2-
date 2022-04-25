package api

import "lesson_4/storage"

// Config Конфиг для нашего приложения
type Config struct {
	BindAddr    string `toml:"bind_addr" env:"BIND_ADDR"`
	LoggerLevel string `toml:"logger_level" env:"LOGGER_LEVEL"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8000",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
