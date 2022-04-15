package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

type Storage struct {
	config *Config
	db     *sql.DB
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
