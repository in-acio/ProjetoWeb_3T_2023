package database

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type VoteDB struct {
	DB *gorm.DB
}

func NewVoteDB(db *gorm.DB) *VoteDB {
	return &VoteDB{
		DB: db,
	}
}

func (v *VoteDB) Create(vote *entity.Vote) error{
	return v.DB.Create(vote).Error
}

func (v *VoteDB) FindByIds(userId, itemId uint) (*entity.Vote, error) {
	var vote entity.Vote

	err := v.DB.First(&vote, "user_id = ? AND item_id = ?",userId, itemId).Error
	
	if err != nil {
		return nil, err
	}

	return &vote, nil
}

func (v *VoteDB) FindById(id uint) (*entity.Vote, error) {
	var vote entity.Vote
	err := v.DB.First(&vote, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &vote, nil
}


func (v *VoteDB) Update(vote *entity.Vote) error {
	_, err := v.FindById(vote.ID)
	if err != nil {
		return err
	}

	return v.DB.Save(vote).Error
}