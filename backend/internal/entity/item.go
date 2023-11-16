package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model `json:"-"`
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Img	string `json:"img"`
	Users []*User `gorm:"many2many:votes;" json:"-"`
}

func NewItem(name, img string) *Item {
	return &Item{
		Name: name,
		Img: img,
	}
}

