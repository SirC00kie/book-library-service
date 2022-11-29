package book

import "context"

type Storage interface {
	Create(ctx context.Context, book Book) (string, error)
	FindAll(ctx context.Context) ([]Book, error)
	FindOne(ctx context.Context, id string) (Book, error)
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id string) error
}
