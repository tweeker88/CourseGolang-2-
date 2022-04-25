package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"lesson_4/internal/app/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "user"
)

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)

	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE login = $1", tableUser)

	var founded bool
	var user models.User

	if err := ur.storage.db.QueryRow(query, login).Scan(&user.Id, &user.Login, &user.Password); err != nil {
		return &user, founded, err
	}

	return &user, founded, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)

	var users = make([]*models.User, 0)

	rows, err := ur.storage.db.Query(query)
	defer rows.Close()

	if err != nil {
		return users, err
	}

	for rows.Next() {
		u := models.User{}
		if err := rows.Scan(&u.Id, &u.Login, &u.Password); err != nil {
			logrus.Info(err)
			continue
		}
		users = append(users, &u)
	}

	return users, nil
}
