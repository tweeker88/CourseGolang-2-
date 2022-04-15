package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"lesson_3/internal/app/api"
	"log"
	"path/filepath"
)

const TOML_FORMAT = ".toml"

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}
func main() {
	flag.Parse()
	config := api.NewConfig()

	var errOpenConfig error
	if filepath.Ext(configPath) == TOML_FORMAT {
		_, errOpenConfig = toml.DecodeFile(configPath, &config)
	} else {
		errOpenConfig = godotenv.Load(configPath)
		errOpenConfig = env.Parse(config)
	}

	if errOpenConfig != nil {
		log.Println("can not find configs file. using default values:", errOpenConfig)
	}

	server := api.NewApi(config)
	logrus.Fatalln(server.Start())
}
