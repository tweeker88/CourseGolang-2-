package api

import (
	"github.com/sirupsen/logrus"
	storage2 "lesson_3/storage"
	"net/http"
)

// конфигурируем логгер
func (a *Api) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)

	if err != nil {
		return err
	}

	a.logger.SetLevel(logLevel)

	return nil
}

// конфигурируем роутер
func (a *Api) configureRouter() {
	a.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I'm server"))
	})
}

func (a *Api) configureStorage() error {
	storage := storage2.NewStorage(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}

	a.storage = storage

	return nil
}