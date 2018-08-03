package book

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/phungvandat/example-go/domain"
	"github.com/phungvandat/example-go/service"
)

// CreateData data for CreateBook
type CreateData struct {
	Name        string      `json:"name"`
	CategoryID  domain.UUID `json:"category_id"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
}

// CreateRequest request struct for CreateBook
type CreateRequest struct {
	Book CreateData `json:"book"`
}

// CreateResponse response struct for CreateBook
type CreateResponse struct {
	Book domain.Book `json:"book"`
}

// StatusCode customstatus code for success create Book
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Book
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			book = &domain.Book{
				Name:        req.Book.Name,
				CategoryID:  req.Book.CategoryID,
				Author:      req.Book.Author,
				Description: req.Book.Description,
			}
		)

		err := s.BookService.Create(ctx, book)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Book: *book}, nil
	}
}

// FindRequest request struct for Find a Book
type FindRequest struct {
	BookID domain.UUID
}

// FindResponse response struct for Find a Book
type FindResponse struct {
	Book *domain.Book `json:"book"`
}

// MakeFindEndPoint make endpoint for find Book
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var bookFind domain.Book
		req := request.(FindRequest)
		bookFind.ID = req.BookID

		book, err := s.BookService.Find(ctx, &bookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Book: book}, nil
	}
}

/*optional*/
//FindByNameRequest request struct for FindByName Book
type FindByNameRequest struct {
	BookName string
}

//FindByNameResponse response struct for FindByName Book
type FindByNameResponse struct {
	Book []domain.Book `json:"books"`
}

//MakeFindByNameEndpoint make endpoint for FindByName Book
func MakeFindByNameEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var bookFindByName domain.Book
		req := request.(FindByNameRequest)
		bookFindByName.Name = req.BookName
		book, err := s.BookService.FindByName(ctx, &bookFindByName)
		if err != nil {
			return nil, err
		}
		return FindByNameResponse{Book: book}, nil
	}
}

//FindByStatusRequset request struct for FindByStatus Book
type FindByStatusRequest struct {
	BookStatus string
}

//FindByStatusResponse response struct for FindByStatus Book
type FindByStatusResponse struct {
	Book []domain.Book `json:"books"`
}

func MakeFindByStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindByStatusRequest)
		BookStatus := req.BookStatus
		book, err := s.BookService.FindByStatus(ctx, BookStatus)
		if err != nil {
			return nil, err
		}
		return FindByStatusResponse{Book: book}, nil
	}
}

//FindByNameAndStatusRequest request struct for FindByNameAndStatus Book
type FindByNameAndStatusRequest struct {
	BookName   string
	BookStatus string
}

//FindByNameAndStatusResponse response struct for FindByNameAndStatus Book
type FindByNameAndStatusRespose struct {
	Book []domain.Book `json:"books"`
}

//MakeFindByNameAndStatusEndpoint make endpoint for FindByNameAndStatus Book
func MakeFindByNameAndStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var bookFindByName domain.Book
		var BookStatus string
		req := request.(FindByNameAndStatusRequest)
		bookFindByName.Name = req.BookName
		BookStatus = req.BookStatus
		book, err := s.BookService.FindByNameAndStatus(ctx, &bookFindByName, BookStatus)
		if err != nil {
			return nil, err
		}
		return FindByNameAndStatusRespose{Book: book}, nil
	}
}

/*optional*/

// FindAllRequest request struct for FindAll Book
type FindAllRequest struct{}

// FindAllResponse request struct for find all Book
type FindAllResponse struct {
	Books []domain.Book `json:"books"`
}

// MakeFindAllEndpoint make endpoint for find all Book
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		books, err := s.BookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Books: books}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID          domain.UUID `json:"-"`
	Name        string      `json:"name"`
	CategoryID  domain.UUID `json:"category_id"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Book UpdateData `json:"book"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Book domain.Book `json:"book"`
}

// MakeUpdateEndpoint make endpoint for update a Book
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			book = domain.Book{
				Model:       domain.Model{ID: req.Book.ID},
				Name:        req.Book.Name,
				CategoryID:  req.Book.CategoryID,
				Author:      req.Book.Author,
				Description: req.Book.Description,
			}
		)

		res, err := s.BookService.Update(ctx, &book)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Book: *res}, nil
	}
}

// DeleteRequest request struct for delete a Book
type DeleteRequest struct {
	BookID domain.UUID
}

// DeleteResponse response struct for Find a Book
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Book
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			bookFind = domain.Book{}
			req      = request.(DeleteRequest)
		)
		bookFind.ID = req.BookID

		err := s.BookService.Delete(ctx, &bookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
