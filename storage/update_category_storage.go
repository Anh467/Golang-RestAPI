package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) UpdateCategory(ctx context.Context, category entities.CategoryUpdate, categoryid int) *entities.CategoryModel {
	// check exitinct of category
	if err := s.db.Select("CategoryID").First(nil).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		}
		panic(err)
	}
	// định nghĩa categoryTemp
	var categoryTemp = &entities.CategoryModel{
		Name: category.Name,
	}
	// tạo ra category
	if err := s.db.Where("CatgoryID = ?", categoryid).Updates(&categoryTemp).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			panic(common.CATEGORY_NAME_DUPLICATED)
		} else {
			panic(err)
		}
	}
	// trả về category vừa tạo
	return categoryTemp
}
