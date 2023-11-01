package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateCategory(ctx context.Context, category entities.CategoryCreate) *entities.CategoryModel {
	// định nghĩa categoryTemp
	var categoryTemp = &entities.CategoryModel{
		Name: category.Name,
	}
	// tạo ra category
	if err := s.db.Create(&categoryTemp).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			panic(common.CATEGORY_NAME_DUPLICATED)
		} else {
			panic(err)
		}
	}
	// trả về category vừa tạo
	return categoryTemp
}
