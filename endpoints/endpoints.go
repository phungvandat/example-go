package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/example-go/service"

	"github.com/phungvandat/example-go/endpoints/book"
	"github.com/phungvandat/example-go/endpoints/category"
	"github.com/phungvandat/example-go/endpoints/lend_book"
	"github.com/phungvandat/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	FindBook                endpoint.Endpoint
	FindByNameBook          endpoint.Endpoint
	FindByStatusBook        endpoint.Endpoint
	FindByNameAndStatusBook endpoint.Endpoint
	FindAllBook             endpoint.Endpoint
	CreateBook              endpoint.Endpoint
	UpdateBook              endpoint.Endpoint
	DeleteBook              endpoint.Endpoint

	FindLendBook    endpoint.Endpoint
	FindAllLendBook endpoint.Endpoint
	CreateLendBook  endpoint.Endpoint
	UpdateLendBook  endpoint.Endpoint
	DeleteLendBook  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		FindBook:                book.MakeFindEndPoint(s),
		FindByNameBook:          book.MakeFindByNameEndpoint(s),
		FindByStatusBook:        book.MakeFindByStatusEndpoint(s),
		FindByNameAndStatusBook: book.MakeFindByNameAndStatusEndpoint(s),
		FindAllBook:             book.MakeFindAllEndpoint(s),
		CreateBook:              book.MakeCreateEndpoint(s),
		UpdateBook:              book.MakeUpdateEndpoint(s),
		DeleteBook:              book.MakeDeleteEndpoint(s),

		FindLendBook:    lend_book.MakeFindEndPoint(s),
		FindAllLendBook: lend_book.MakeFindAllEndpoint(s),
		CreateLendBook:  lend_book.MakeCreateEndpoint(s),
		UpdateLendBook:  lend_book.MakeUpdateEndpoint(s),
		DeleteLendBook:  lend_book.MakeDeleteEndpoint(s),
	}
}
