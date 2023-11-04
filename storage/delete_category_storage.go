package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) DeleteCategory(ctx context.Context, categoryid int) {
	// check exitinct of category
	if err := s.db.Table(entities.CategorieModelTable).Where("CategoryID = ?", categoryid).Take(&entities.CategoryModel{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		}
		panic(err)
	}
	// delete that categoryid
	if err := s.db.Table(entities.CategorieModelTable).Where("CategoryID = ?", categoryid).Table(entities.CategorieModelTable).Delete(nil).Error; err != nil {
		panic(err)
	}
}
