package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound       = errNotFound{}
	ErrUnknown        = errUnknown{}
	ErrRecordNotFound = errRecordNotFound{}
	ErrNameIsRequired = errNameIsRequired{}
	ErrminimumLength  = errMinimumLength{}
	ErrExistName      = errExistName{}
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

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "Category name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errMinimumLength struct{}

func (errMinimumLength) Error() string {
	return "Name of category is length > 5 characters"
}
func (errMinimumLength) StatusCode() int {
	return http.StatusBadRequest
}

type errExistName struct{}

func (errExistName) Error() string {
	return "Name is exist in database"
}
func (errExistName) StatusCode() int {
	return http.StatusBadRequest
}
