package services

import "github.com/ascribner/sustainabilityapp/internal/repos"

type ItemService interface {
}

type ItemServiceImpl struct {
	ir repos.ItemRepo
}

func NewItemService(ir repos.ItemRepo) ItemService {
	return &ItemServiceImpl{
		ir: ir,
	}
}
