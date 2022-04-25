package storage

type Config struct {
	DataBaseURI string `toml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}
