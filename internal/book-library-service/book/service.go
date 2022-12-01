package book

import (
	"context"
	"fmt"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (string, error) {
	id, err := s.repository.Create(ctx, Book{
		ID:          "",
		Name:        dto.Name,
		Type:        dto.Type,
		Content:     dto.Content,
		Author:      dto.Author,
		Year:        dto.Year,
		Description: dto.Description,
	})
	if err != nil {
		return "", fmt.Errorf("create book error: %v", err)
	}
	return id, nil
}

func (s *Service) FindAll(ctx context.Context) (books []Book, err error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("find all book error: %v", err)
	}
	return all, nil
}

func (s *Service) FindOne(ctx context.Context, id string) (b Book, err error) {
	one, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return Book{}, fmt.Errorf("find one book error: %v", err)
	}
	return one, nil
}

func (s *Service) Update(ctx context.Context, book Book) error {
	err := s.repository.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("update one book error: %v", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete one user error: %v", err)
	}
	return nil
}
