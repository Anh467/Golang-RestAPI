package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) DeleteCategory(ctx context.Context, categoryid string) {
	// check exitinct of category
	if err := s.db.Select("CategoryID").Where("CategoryID = ?", categoryid).First(nil).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CATEGORY_ID_NOT_EXIST)
		}
		panic(err)
	}
	// delete that categoryid
	if err := s.db.Table(entities.CategorieModelTable).Where("CategoryID = ?", categoryid).Delete(nil).Error; err != nil {
		panic(err)
	}
}
