package api

import (
	"github.com/sirupsen/logrus"
	storage2 "lesson_4/storage"
	"net/http"
)

var (
	prefix string = "/api/v1"
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
	a.logger.Info("Routes configuration")
	a.router.HandleFunc(prefix+"/article", a.GetArticles).Methods(http.MethodGet)
	a.router.HandleFunc(prefix+"/article/{id}", a.UpdateArticle).Methods(http.MethodPut)
	a.router.HandleFunc(prefix+"/article/{id}", a.GetArticle).Methods(http.MethodGet)
	a.router.HandleFunc(prefix+"/article/{id}", a.DeleteArticle).Methods(http.MethodDelete)
	a.router.HandleFunc(prefix+"/article", a.CreateArticle).Methods(http.MethodPost)
	a.router.HandleFunc(prefix+"/user/register", a.PostRegisterUser).Methods(http.MethodPost)

}

func (a *Api) configureStorage() error {
	storage := storage2.NewStorage(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}

	a.storage = storage

	return nil
}
