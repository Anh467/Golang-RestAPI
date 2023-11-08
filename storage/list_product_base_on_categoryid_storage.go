package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) ListProductBaseOnCategoryIDStorage(ctx context.Context, categoryid, offsetNum int, limitNum int) []entities.ProductModel {
	var products []entities.ProductModel
	var count int64
	// check exitinct of category
	if err := s.db.Table(entities.CategorieModelTable).Where("CategoryID = ?", categoryid).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(common.CATEGORY_ID_NOT_EXIST)
	}
	// get data
	if err := s.db.Table(entities.ProductModelTable).
		Offset(offsetNum*limitNum).
		Limit(limitNum).
		Where("CategoryID = ?", categoryid).
		Preload("Category").
		Find(&products).Error; err != nil {
		panic(err)
	}
	// return
	return products
}
