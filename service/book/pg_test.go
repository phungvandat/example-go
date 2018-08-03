package book

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"

	testutil "github.com/phungvandat/example-go/config/database/pg/util"
	"github.com/phungvandat/example-go/domain"
)

func TestPGService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Book{
					Name:        "why not love me.",
					CategoryID:  category.ID,
					Author:      "Phung van dat",
					Description: "the book is very bad",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_Update(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success update",
			args: args{
				&domain.Book{
					Model:       domain.Model{ID: book.ID},
					Name:        "why not love me.",
					CategoryID:  category.ID,
					Author:      "Phung van dat",
					Description: "the book is very bad",
				},
			},
		},
		{
			name: "failed update",
			args: args{
				&domain.Book{
					Model:       domain.Model{ID: fakeBookID},
					Name:        "why not love me.",
					CategoryID:  category.ID,
					Author:      "Phung van dat",
					Description: "the book is very bad",
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			_, err := s.Update(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_Find(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Book
		wantErr error
	}{
		{
			name: "success find correct book",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: book.ID},
				},
			},
			want: &domain.Book{
				Model: domain.Model{ID: book.ID},
			},
		},
		{
			name: "failed find book by not exist book id",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: fakeBookID},
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}

			got, err := s.Find(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && got.ID.String() != tt.want.ID.String() {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindByName(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindByName(context.Background(), tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindByStatus(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		Status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindByStatus(context.Background(), tt.args.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindByNameAndStatus(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		p      *domain.Book
		Status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindByNameAndStatus(context.Background(), tt.args.p, tt.args.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindByNameAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindByNameAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_Delete(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success delete",
			args: args{
				&domain.Book{
					Model:       domain.Model{ID: book.ID},
					Name:        "why not love me.",
					CategoryID:  domain.MustGetUUIDFromString("dc9076e9-2fda-4019-bd2c-900a8284b9c4"),
					Author:      "Phung van dat",
					Description: "the book is very bad",
				},
			},
		},
		{
			name: "failed delete by not exist book id",
			args: args{
				&domain.Book{
					Model:       domain.Model{ID: fakeBookID},
					Name:        "why not love me.",
					CategoryID:  domain.MustGetUUIDFromString("dc9076e9-2fda-4019-bd2c-900a8284b9c4"),
					Author:      "Phung van dat",
					Description: "the book is very bad",
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			err := s.Delete(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
