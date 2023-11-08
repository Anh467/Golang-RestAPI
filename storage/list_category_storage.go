package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListCategoryStorage(ctx context.Context, offset, limit int) []entities.CategoryModel {
	var categories []entities.CategoryModel
	if err := s.db.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	return categories
}
