package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string `json:"name"`
	Img	string `json:"img"`
	Users []*User `gorm:"many2many:votes;"`
}

func NewItem(name, img string) *Item {
	return &Item{
		Name: name,
		Img: img,
	}
}

