package storage

import (
	"common"
	"context"
	"entities"
)

func (s *sqlserverStore) CreateProduct(ctx context.Context, product *entities.Product) *entities.Product {

	var category *entities.CategorieModel
	// check exist category
	if err := s.db.Where(" CategoryID = ? ", product.CategoryID).First(&category).Error; err != nil {
		panic(common.CATEGORY_ID_NOT_EXIST)
	}
	if category == nil {
		panic(common.CATEGORY_ID_NOT_EXIST)
	}
	// create product
	if err := s.db.Create(&product).Error; err != nil {
		panic(err)
	}
	// check valid product
	if product == nil {
		panic(common.PRODUCT_CREATE)
	}
	product.Category = *category
	return product
}
