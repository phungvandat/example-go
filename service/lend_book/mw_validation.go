package lend_book

import (
	"context"

	"github.com/phungvandat/example-go/domain"
)

// Declare Regex
type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, lend_book *domain.LendBook) (err error) {
	if lend_book.BookID.IsZero() {
		return ErrBookIDIsRequired
	}

	if lend_book.UserID.IsZero() {
		return ErrUserIDIsRequired
	}
	if lend_book.From.IsZero() {
		return ErrFromIsRequired
	}
	if lend_book.To.IsZero() {
		return ErrToIsRequired
	}

	return mw.Service.Create(ctx, lend_book)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.LendBook, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lend_book *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Find(ctx, lend_book)
}

func (mw validationMiddleware) Update(ctx context.Context, lend_book *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Update(ctx, lend_book)
}
func (mw validationMiddleware) Delete(ctx context.Context, lend_book *domain.LendBook) error {
	return mw.Service.Delete(ctx, lend_book)
}
