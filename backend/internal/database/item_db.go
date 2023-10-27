package database

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type ItemDB struct {
	DB *gorm.DB
}

func NewItemDB(db *gorm.DB) *ItemDB {
	return &ItemDB{
		DB: db,
	}
}

func (i *ItemDB) Create(item *entity.Item) error {
	return i.DB.Create(item).Error
}

func (i *ItemDB) FindById(id uint) (*entity.Item, error) {
	var item entity.Item
	err := i.DB.First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *ItemDB) FindAll() ([]entity.Item, error) {
	var itens []entity.Item
	err := i.DB.Find(&itens).Error

	if err != nil {
		return nil, err
	}

	return itens, nil
}

func (i *ItemDB) Update(item *entity.Item) error {
	_, err := i.FindById(item.ID)
	if err != nil {
		return err
	}

	return i.DB.Save(item).Error
}

func (i *ItemDB) Delete(id uint) error {
	item, err := i.FindById(id)
	if err != nil {
		return err
	}

	return i.DB.Delete(item).Error
}