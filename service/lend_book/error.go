package lend_book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrRecordNotFound   = errRecordNotFound{}
	ErrBookIDIsRequired = errBookIDIsRequired{}
	ErrUserIDIsRequired = errUserIDIsRequired{}
	ErrFromIsRequired   = errFromIsRequired{}
	ErrToIsRequired     = errToIsRequired{}
	ErrBookIDNotExist   = errBookIDNotExist{}
	ErrUserIDNotExist   = errUserIDNotExist{}
	ErrLendedBook       = errLendedBook{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errBookIDIsRequired struct{}

func (errBookIDIsRequired) Error() string {
	return "ID of book is required"
}
func (errBookIDIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errUserIDIsRequired struct{}

func (errUserIDIsRequired) Error() string {
	return "ID of user is required"
}
func (errUserIDIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errFromIsRequired struct{}

func (errFromIsRequired) Error() string {
	return "From is required"
}
func (errFromIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errToIsRequired struct{}

func (errToIsRequired) Error() string {
	return "To is required"
}
func (errToIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIDNotExist struct{}

func (errBookIDNotExist) Error() string {
	return "ID of book not exist in table books"
}
func (errBookIDNotExist) StatusCode() int {
	return http.StatusNotFound
}

type errUserIDNotExist struct{}

func (errUserIDNotExist) Error() string {
	return "ID of user not exist in table users"
}
func (errUserIDNotExist) StatusCode() int {
	return http.StatusNotFound
}

type errLendedBook struct{}

func (errLendedBook) Error() string {
	return "The book is lended"
}
func (errLendedBook) StatusCode() int {
	return http.StatusBadRequest
}
