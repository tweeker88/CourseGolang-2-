package storage

type Config struct {
	DataBaseURI string `toml:"database" env:"DATABASE"`
}

func NewConfig() *Config {
	return &Config{}
}
