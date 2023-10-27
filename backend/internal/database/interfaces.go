package database

import "backend/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindById(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
}

type ItemInterface interface {
	Create(item *entity.Item) error
	FindById(id uint) (*entity.Item, error)
	FindAll() ([]entity.Item, error) 
	Update(item *entity.Item) error
	Delete(id uint) error
}

type VoteInterface interface {
	Create(vote *entity.Vote) error
	FindById(id uint) (*entity.Vote, error)
	FindByIds(userId, itemId uint) (*entity.Vote, error)
	Update(vote *entity.Vote) error
}