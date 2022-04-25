package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func NewStorage(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DataBaseURI)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

func (storage *Storage) Close() error {
	err := storage.db.Close()

	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) User() *UserRepository {
	if storage.userRepository == nil {
		storage.userRepository = &UserRepository{
			storage: storage,
		}
	}

	return storage.userRepository
}

func (storage *Storage) Article() *ArticleRepository {
	if storage.articleRepository == nil {
		storage.articleRepository = &ArticleRepository{
			storage: storage,
		}
	}

	return storage.articleRepository
}
