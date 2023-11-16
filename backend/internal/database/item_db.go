package database

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type ItemDB struct {
	DB *gorm.DB
}

type TopItem struct {
    entity.Item
    VoteCount uint `json:"vote_count"`
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

func (i *ItemDB) FindAll(page, limit int, sort string) ([]entity.Item, error) {
	var items []entity.Item
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = i.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&items).Error
	} else {
		err = i.DB.Order("created_at " + sort).Find(&items).Error
	}

	return items, err
}

func (i *ItemDB) GetRanking() ([][]TopItem, error){
	var ranking [][]TopItem

	for value := 0; value <= 2; value++ {
		var topItems []TopItem

		i.DB.Table("items").
			Select("items.*, COUNT(votes.id) as vote_count").
			Joins("JOIN votes ON items.id = votes.item_id").
			Where("votes.value = ?", value).
			Group("items.id").
			Order("vote_count DESC").
			Limit(5).
			Find(&topItems)

			

		ranking = append(ranking, topItems)
	}

	return ranking, nil
}

func (i *ItemDB) Update(item *entity.Item) error {
	_, err := i.FindById(item.ID)
	if err != nil {
		return err
	}

	return i.DB.Save(item).Error
}

func (i *ItemDB) Delete(id uint) error {
	err := i.DB.Unscoped().Exec("DELETE FROM votes WHERE item_id = ?", id).Error
	if err != nil {
		return err
	}
	
	item, err := i.FindById(id)
	if err != nil {
		return err
	}

	return i.DB.Delete(item).Error
}