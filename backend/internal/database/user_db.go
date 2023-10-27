package database

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{ DB: db }
}

func (u *UserDB) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindById(id uint) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDB) Update(user *entity.User) error {
	_, err := u.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	return u.DB.Save(user).Error
}