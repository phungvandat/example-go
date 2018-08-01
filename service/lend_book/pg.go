package lend_book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phungvandat/example-go/domain"
)

// pgService implmenter for LendBook serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for LendBook service
func (s *pgService) Create(_ context.Context, p *domain.LendBook) error {
	var errExistBoID = s.db.Where("id = ?", p.BookID).Find(&domain.Book{}).Error
	if errExistBoID != nil {
		if errExistBoID == gorm.ErrRecordNotFound {
			return ErrBookIDNotExist
		}
		return errExistBoID
	}
	var errLendedBo = s.db.Where("book_id  = ?", p.BookID).Find(&domain.LendBook{}).Error
	if errLendedBo == nil {
		return ErrLendedBook
	}
	var errExistUsID = s.db.Where("id = ?", p.UserID).Find(&domain.User{}).Error
	if errExistUsID != nil {
		if errExistUsID == gorm.ErrRecordNotFound {
			return ErrUserIDNotExist
		}
		return errExistUsID
	}
	return s.db.Create(p).Error
}

// Update implement Update for LendBook service
func (s *pgService) Update(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if !p.BookID.IsZero() {
		if p.BookID != old.BookID {
			var errExistBoID = s.db.Where("id = ?", p.BookID).Find(&domain.Book{}).Error
			if errExistBoID != nil {
				if errExistBoID == gorm.ErrRecordNotFound {
					return nil, ErrBookIDNotExist
				}
				return nil, errExistBoID
			}
			var errLendedBo = s.db.Where("book_id  = ?", p.BookID).Find(&domain.LendBook{}).Error
			if errLendedBo == nil {
				return nil, ErrLendedBook
			}
			old.BookID = p.BookID
		}
	}
	if !p.UserID.IsZero() {
		if p.UserID != old.UserID {
			var errExistUsID = s.db.Where("id = ?", p.UserID).Find(&domain.User{}).Error
			if errExistUsID != nil {
				if errExistUsID == gorm.ErrRecordNotFound {
					return nil, ErrUserIDNotExist
				}
				return nil, errExistUsID
			}
			old.UserID = p.UserID
		}
	}
	if !p.From.IsZero() {
		old.From = p.From
	}
	if !p.To.IsZero() {
		old.To = p.To
	}
	return &old, s.db.Save(&old).Error
}

// Find implement Find for LendBook service
func (s *pgService) Find(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendBook service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendBook, error) {
	res := []domain.LendBook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendBook service
func (s *pgService) Delete(_ context.Context, p *domain.LendBook) error {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
