package storage

import (
	"common"
	"context"
	"entities"
)

func (s *sqlserverStore) GetProduct(ctx context.Context, productid string) *entities.Product {
	var product *entities.Product
	if err := s.db.Where("ProductID = ?", productid).First(&product).Error; err != nil {
		panic(err)
	}
	if err := s.db.Where("CategoryID = ?", &product.CategoryID).First(&product.Category).Error; err != nil {
		panic(err)
	}
	if product == nil {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
	return product
}
