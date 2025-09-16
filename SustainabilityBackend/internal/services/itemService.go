package services

import (
	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/repos"
)

type ItemService interface {
	AddItem(entity.ItemEntity, string) error
	GetItems(string) ([]entity.Item, error)
}

type ItemServiceImpl struct {
	ir repos.ItemRepo
}

func NewItemService(ir repos.ItemRepo) ItemService {
	return &ItemServiceImpl{
		ir: ir,
	}
}

func (is *ItemServiceImpl) AddItem(item entity.ItemEntity, user_email string) error {
	if err := is.ir.InsertItem(item.GetTitle(), item.GetOffset(), item.GetDate(), user_email); err != nil {
		return err
	}

	return nil
}

func (is *ItemServiceImpl) GetItems(user_email string) ([]entity.Item, error) {
	items, err := is.ir.GetItems(user_email)
	if err != nil {
		return nil, err
	}

	return items, nil
}
