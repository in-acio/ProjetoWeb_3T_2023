package entity

import (
	"errors"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrIDIsRequired = errors.New("ID is required")
	ErrNameIsRequired = errors.New("Name is required")
	ErrEmailIsRequired = errors.New("Email is required")
	ErrPasswordIsRequired = errors.New("Password is invalid")
)

type User struct {
	gorm.Model `json:"-"`
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin bool `json:"isAdmin"`
	Items []*Item `gorm:"many2many:votes;"`
}

func NewUser(name, email, password string) (*User, error) { 
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &User{
		Name: name,
		Email: email,
		Password: string(hash),
		IsAdmin: false,
	}

	err = newUser.Validate()
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) Validate() error {
	emailAddress, err := mail.ParseAddress(u.Email)
    if err != nil {
		return err
	}

	if (strings.Split(emailAddress.Address, "@")[1]) != "aluno.feliz.ifrs.edu.br" {
		return errors.New("email inválido")
	}

	return nil;
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}