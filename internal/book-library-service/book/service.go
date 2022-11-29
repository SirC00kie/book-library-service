package book

import (
	"book-library-service/pkg/logging"
	"context"
	"fmt"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (b Book, err error) {
	_, err = s.storage.Create(ctx, Book{
		ID:          "",
		Name:        dto.Name,
		Type:        dto.Type,
		Content:     dto.Content,
		Author:      dto.Author,
		Year:        dto.Year,
		Description: dto.Description,
	})
	if err != nil {
		return Book{}, fmt.Errorf("create book error: %v", err)
	}
	return
}

func (s *Service) FindAll(ctx context.Context) (books []Book, err error) {
	all, err := s.storage.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("find all book error: %v", err)
	}
	return all, nil
}

func (s *Service) FindOne(ctx context.Context, id string) (b Book, err error) {
	one, err := s.storage.FindOne(ctx, id)
	if err != nil {
		return Book{}, fmt.Errorf("find one book error: %v", err)
	}
	return one, nil
}

func (s *Service) Update(ctx context.Context, book Book) error {
	err := s.storage.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("update one book error: %v", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.storage.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete one user error: %v", err)
	}
	return nil
}
