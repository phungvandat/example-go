package book

import (
	"context"

	"github.com/phungvandat/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Book) error
	Update(ctx context.Context, p *domain.Book) (*domain.Book, error)
	Find(ctx context.Context, p *domain.Book) (*domain.Book, error)
	FindByName(ctx context.Context, p *domain.Book) ([]domain.Book, error)
	FindByStatus(ctx context.Context, status string) ([]domain.Book, error)
	FindByNameAndStatus(ctx context.Context, p *domain.Book, status string) ([]domain.Book, error)
	FindAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, p *domain.Book) error
}
