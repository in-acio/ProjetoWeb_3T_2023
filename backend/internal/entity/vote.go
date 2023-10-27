package entity

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID  uint
	ItemID  uint
	Value   uint
}

func NewVote(user_id, item_id, value uint) *Vote {
	return &Vote{
		UserID: user_id,
		ItemID: item_id,
		Value: value,
	}
}

