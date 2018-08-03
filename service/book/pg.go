package book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phungvandat/example-go/domain"
)

// pgService implmenter for Book serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Book service
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	// Check id of category exist in table categories
	var checkErr = s.db.Where("id = ?", p.CategoryID).Find(&domain.Category{}).Error
	if checkErr != nil {
		if checkErr == gorm.ErrRecordNotFound {
			return ErrNotExistCategoryID
		}
		return checkErr
	}
	return s.db.Create(p).Error
}

// Update implement Update for Book service
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if p.Name != "" {
		old.Name = p.Name
	}
	if !p.CategoryID.IsZero() {
		old.CategoryID = p.CategoryID
		// Check id of category exist in table categories
		var checkExist = s.db.Where("id = ?", p.CategoryID).Find(&domain.Category{}).Error
		if checkExist != nil {
			if checkExist == gorm.ErrRecordNotFound {
				return nil, ErrNotExistCategoryID
			}
			return nil, checkExist
		}
	}

	if p.Author != "" {
		old.Author = p.Author
	}
	if p.Description != "" {
		old.Description = p.Description
	}

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Book service
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

//FindByName implement FindByName for Book service
func (s *pgService) FindByName(_ context.Context, p *domain.Book) ([]domain.Book, error) {
	res := []domain.Book{}
	if err := s.db.Where("name = ?", p.Name).Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

//FindByStatus implement FindByStatus for Book service
func (s *pgService) FindByStatus(_ context.Context, status string) ([]domain.Book, error) {
	res := []domain.Book{}
	if status == "0" {
		if err := s.db.Not("id", s.db.Select("book_id").Find(&domain.LendBook{}).QueryExpr()).Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ErrNotFound
			}
			return nil, err
		}
	} else {
		if err := s.db.Joins("join LEND_BOOKS on LEND_BOOKS.book_id = BOOKS.id").Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ErrNotFound
			}
			return nil, err
		}
	}
	return res, nil
}

//FindByNameAndStatus implement FindByNameAndStatus for Book service
func (s *pgService) FindByNameAndStatus(_ context.Context, p *domain.Book, status string) ([]domain.Book, error) {
	res := []domain.Book{}
	if status == "1" {
		if err := s.db.Joins("join LEND_BOOKS on LEND_BOOKS.book_id = BOOKS.id").Where("BOOKS.name = ? ", p.Name).Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ErrNotFound
			}
			return nil, err
		}
	} else {
		if err := s.db.Not("id", s.db.Select("book_id").Find(&domain.LendBook{}).QueryExpr()).Where("name = ?", p.Name).Find(&res).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err == gorm.ErrRecordNotFound {
					return nil, ErrNotFound
				}
				return nil, err
			}
		}
	}
	return res, nil
}

// FindAll implement FindAll for Book service
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Book service
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
