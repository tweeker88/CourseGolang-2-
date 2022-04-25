package storage

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"lesson_4/internal/app/models"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "article"
)

func (ar *ArticleRepository) Create(a models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)

	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.Id); err != nil {
		return nil, err
	}

	return &a, nil
}

func (ar *ArticleRepository) FindById(id int) (*models.Article, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", tableArticle)

	var founded bool
	var article models.Article

	if err := ar.storage.db.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.Author, &article.Content); err != nil {
		return &article, founded, err
	}

	return &article, true, nil

}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)

	var articles = make([]*models.Article, 0)

	rows, err := ar.storage.db.Query(query)
	defer rows.Close()

	if err != nil {
		return articles, err
	}

	for rows.Next() {
		a := models.Article{}
		if err := rows.Scan(&a.Id, &a.Title, &a.Author, &a.Content); err != nil {
			logrus.Info(err)
			continue
		}
		articles = append(articles, &a)
	}

	return articles, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	foundedArticle, founded, err := ar.FindById(id)
	if founded == false {
		err := errors.New(fmt.Sprintf("Article with id %v not found", id))
		logrus.Errorln(err)
		return nil, err
	}

	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableArticle)
	_, err = ar.storage.db.Exec(query, id)

	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return foundedArticle, nil
}

func (ar *ArticleRepository) Update(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("UPDATE %s SET title = $1, author = $2, content = $3  WHERE id = $4", tableArticle)
	_, err := ar.storage.db.Exec(query, article.Title, article.Author, article.Content, article.Id)

	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return article, nil
}
