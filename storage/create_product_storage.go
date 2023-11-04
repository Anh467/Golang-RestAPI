package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) CreateProduct(ctx context.Context, product *entities.ProductModel) *entities.ProductModel {

	var category *entities.CategoryModel
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
