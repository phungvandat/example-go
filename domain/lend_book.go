package domain

import (
	"time"
)

//Lend book describe lend book in system
type LendBook struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"time"`
}
