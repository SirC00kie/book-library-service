package book

import (
	"book-library-service/internal/book-library-service/apperror"
	"book-library-service/internal/book-library-service/handlers"
	"book-library-service/pkg/logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/books"
	userURL  = "/books/:uuid"
)

type handler struct {
	logger      *logging.Logger
	bookService *Service
}

func NewHandler(logger *logging.Logger, service *Service) handlers.Handler {
	return &handler{
		logger:      logger,
		bookService: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartialUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(writer http.ResponseWriter, request *http.Request) error {
	b, err := h.bookService.FindAll(context.Background())
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed get list: %v", err)
		writer.Write([]byte("this is list of users"))
	}
	book, err := json.Marshal(b)
	writer.WriteHeader(200)
	writer.Write([]byte(book))

	return nil
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(201)
	writer.Write([]byte("this is create book"))

	return nil
}

func (h *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(200)
	writer.Write([]byte("this is book by uuid"))

	return nil
}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is update book"))

	return nil
}

func (h *handler) PartialUpdateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is partial update book"))

	return nil
}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is delete book"))

	return nil
}
