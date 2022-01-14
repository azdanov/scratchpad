package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Scratches interface {
	Insert(string, string, string) (int, error)
	Get(int) (*Scratch, error)
	Latest() ([]*Scratch, error)
}

type Scratch struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type Users interface {
	Insert(string, string, string) error
	Authenticate(string, string) (int, error)
	Get(int) (*User, error)
	ChangePassword(int, string, string) error
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
