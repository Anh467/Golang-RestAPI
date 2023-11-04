package storage

import (
	"context"
	"main/entities"
)

func (s *sqlserverStore) ListProduct(ctx context.Context, offsetNum int, limitNum int) ([]entities.ProductModel, error) {
	var products []entities.ProductModel
	// connect 2 table Product and Catgory together
	// get data
	if err := s.db.Table(entities.ProductModelTable).Offset(offsetNum * limitNum).Limit(limitNum).Find(&products).Error; err != nil {
		return nil, err
	}
	// get Categories
	for index, element := range products {
		s.db.Table(entities.CategorieModelTable).Where("CategoryID = ?", element.CategoryID).Find(&products[index].Category)
	}
	// return
	return products, nil
}
