package services

import "github.com/ascribner/sustainabilityapp/internal/entity"

func TotalItemOffset(items []entity.Item) float64 {
	var total float64 = 0.0

	for _, item := range items {
		total += item.GetOffset()
	}

	return total
}
