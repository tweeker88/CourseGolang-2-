package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"lesson_3/storage"
	"net/http"
)

// Api Структура главной апишки
type Api struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

// NewApi конструктор для создания нашей апишки
func NewApi(config *Config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start Метод для запуска нашего сервера
func (a *Api) Start() error {
	errConfigureLogger := a.configureLoggerField()

	if errConfigureLogger != nil {
		return errConfigureLogger
	}

	a.logger.Println("Hay! Server is started. Port:", a.config.BindAddr)
	a.configureRouter()

	if errConfigureStorage := a.configureStorage(); errConfigureStorage != nil {
		return errConfigureStorage
	}

	return http.ListenAndServe(a.config.BindAddr, a.router)
}
