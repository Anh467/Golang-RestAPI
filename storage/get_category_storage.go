package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) GetCategory(ctx context.Context, categoryid string) entities.CategoryModel {
	var category entities.CategoryModel
	if err := s.db.Select("CategoryID", "Name").Where("CategoryID = ?", categoryid).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	return category
}
