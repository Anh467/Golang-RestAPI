package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) UpdateProduct(ctx context.Context, product entities.ProductUpdate, productid int) *entities.Product {
	// for biz ProductID bắt buộc phải có còn lại không có cũng được
	// declare product
	var productTemp entities.Product
	// check appropriate category id
	category := entities.CategoryModel{}
	if err := s.db.Where("CategoryID = ?", product.CategoryID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	// get product id
	if err := s.db.Where("ProductID = ?", productid).Model(&productTemp).Updates(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	// get updated data
	if err := s.db.Where("ProductID = ?", productid).First(&productTemp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	productTemp.Category = category
	return &productTemp
}
