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
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateBook))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetBookByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateBook))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteBook))
}

func (h *handler) GetList(writer http.ResponseWriter, request *http.Request) error {
	b, err := h.bookService.FindAll(context.Background())
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed get list books: %v", err)
	}
	book, err := json.Marshal(b)
	writer.WriteHeader(200)
	writer.Write(book)

	return nil
}

func (h *handler) CreateBook(writer http.ResponseWriter, request *http.Request) error {
	req := &CreateUserDTO{}
	err := json.NewDecoder(request.Body).Decode(req)

	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed create book: %v", err)
	}

	b, err := h.bookService.Create(context.Background(), CreateUserDTO{
		Name:        req.Name,
		Type:        req.Type,
		Content:     req.Content,
		Author:      req.Author,
		Year:        req.Year,
		Description: req.Description,
	})
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed create book: %v", err)
	}
	bookId, err := json.Marshal(b)
	writer.WriteHeader(201)
	writer.Write(bookId)

	return nil
}

func (h *handler) GetBookByUUID(writer http.ResponseWriter, request *http.Request) error {
	uuid := request.URL.Query().Get("uuid")
	if uuid == "" {
		writer.WriteHeader(400)
		return fmt.Errorf("failed get book by uuid, %v", uuid)
	}
	b, err := h.bookService.FindOne(context.Background(), uuid)
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed find book by uuid: %v", err)
	}
	book, err := json.Marshal(b)
	writer.WriteHeader(200)
	writer.Write(book)

	return nil
}

func (h *handler) UpdateBook(writer http.ResponseWriter, request *http.Request) error {
	err := h.bookService.Update(context.Background(), Book{
		ID:          "6388dc0ceba6baffd6e6d897",
		Name:        "111",
		Type:        "222",
		Content:     "333",
		Author:      "444",
		Year:        2222,
		Description: "5555",
	})
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed update book: %v", err)
	}
	writer.WriteHeader(204)
	writer.Write([]byte("this is update book"))

	return nil
}

func (h *handler) DeleteBook(writer http.ResponseWriter, request *http.Request) error {
	uuid := request.URL.Query().Get("uuid")
	if uuid == "" {
		writer.WriteHeader(400)
		return fmt.Errorf("failed delete book, %v", uuid)
	}
	err := h.bookService.Delete(context.Background(), uuid)
	if err != nil {
		writer.WriteHeader(400)
		return fmt.Errorf("failed delete book, %v", err)
	}
	writer.WriteHeader(204)
	writer.Write([]byte("this is delete book"))

	return nil
}
