package book

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

func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {
	//check empty and length of name, description
	if book.Name == "" {
		return ErrNameIsRequired
	}
	if len(book.Name) < 5 {
		return ErrMinimumLengthName
	}
	if book.Description == "" {
		return ErrDescriptionIsRequired
	}
	if len(book.Description) < 5 {
		return ErrMinimumLengthDescription
	}
	//check empty of category_id
	if book.CategoryID.IsZero() {
		return ErrCategoryIDIsRequired
	}
	return mw.Service.Create(ctx, book)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}

func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	//check length of name
	if book.Name != "" && len(book.Name) < 5 {
		return nil, ErrMinimumLengthName
	}
	//check length of description
	if book.Description != "" && len(book.Description) < 5 {
		return nil, ErrMinimumLengthDescription
	}
	return mw.Service.Update(ctx, book)
}
func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
