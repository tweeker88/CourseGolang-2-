package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lesson_4/internal/app/models"
	"net/http"
	"strconv"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (a *Api) GetArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	articles, err := a.storage.Article().SelectAll()

	if err != nil {
		a.logger.Errorln(err)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Get All Articles GET /articles")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (a *Api) UpdateArticle(writer http.ResponseWriter, request *http.Request) {
	var articleFromRequest models.Article
	errFromDecode := json.NewDecoder(request.Body).Decode(&articleFromRequest)

	if errFromDecode != nil {
		a.logger.Errorln(errFromDecode)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	idArticle, errAtoi := strconv.Atoi(mux.Vars(request)["id"])

	if errAtoi != nil {
		a.logger.Errorln("Troubles while parsing {id} param:", errAtoi)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	foundedArticle, found, err := a.storage.Article().FindById(idArticle)

	if found == false {
		a.logger.Errorln("Article not found")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Article not found",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err != nil {
		a.logger.Errorln(err)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	articleFromRequest.Id = foundedArticle.Id
	updatedArticle, errForUpdate := a.storage.Article().Update(&articleFromRequest)

	if errForUpdate != nil {
		a.logger.Errorln(errForUpdate)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Update Articles GET /articles/{id}")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedArticle)
}

func (a *Api) GetArticle(writer http.ResponseWriter, request *http.Request) {

	idArticle, errAtoi := strconv.Atoi(mux.Vars(request)["id"])

	if errAtoi != nil {
		a.logger.Errorln("Troubles while parsing {id} param:", errAtoi)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	foundedArticle, found, err := a.storage.Article().FindById(idArticle)

	if found == false {
		a.logger.Errorln("Article not found")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Article not found",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err != nil {
		a.logger.Errorln(err)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Get Articles GET /articles/{id}")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(foundedArticle)
}

func (a *Api) DeleteArticle(writer http.ResponseWriter, request *http.Request) {

	idArticle, errAtoi := strconv.Atoi(mux.Vars(request)["id"])

	if errAtoi != nil {
		a.logger.Errorln("Troubles while parsing {id} param:", errAtoi)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, found, err := a.storage.Article().FindById(idArticle)

	if found == false {
		a.logger.Errorln("Article not found")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Article not found",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err != nil {
		a.logger.Errorln(err)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	deletedArticle, errDeleted := a.storage.Article().DeleteById(idArticle)

	if errDeleted != nil {
		a.logger.Errorln(err)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Delete Article GET /articles/{id}")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(deletedArticle)
}

func (a *Api) CreateArticle(writer http.ResponseWriter, request *http.Request) {
	var articleFromRequest models.Article
	errFromDecode := json.NewDecoder(request.Body).Decode(&articleFromRequest)

	if errFromDecode != nil {
		a.logger.Errorln(errFromDecode)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	createdArticle, errCreated := a.storage.Article().Create(articleFromRequest)

	if errCreated != nil {
		a.logger.Errorln(errCreated)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Create Article POST /articles")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(createdArticle)
}

func (a *Api) PostRegisterUser(writer http.ResponseWriter, request *http.Request) {
	var userFromRequest models.User
	errFromDecode := json.NewDecoder(request.Body).Decode(&userFromRequest)

	if errFromDecode != nil {
		a.logger.Errorln(errFromDecode)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	_, found, _ := a.storage.User().FindByLogin(userFromRequest.Login)

	if found {
		a.logger.Errorln(errFromDecode)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "User with that login already exists in database",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	createdUser, errCreated := a.storage.User().Create(&userFromRequest)

	if errCreated != nil {
		a.logger.Errorln(errCreated)
		msg := Message{
			http.StatusNotImplemented,
			"We have some troubles to accessing articles in database. Try later",
			true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a.logger.Info("Register User POST /user/register")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(createdUser)
}
