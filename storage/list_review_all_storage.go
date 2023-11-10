package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListReviewAllStorage(ctx context.Context, offsetNum int, limitNum int) []entities.ReviewGet {
	// declare list
	var reviews []entities.ReviewGet
	// get list
	if err := s.db.Offset(offsetNum).Limit(limitNum).Preload("User").Find(&reviews).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.REVIEW_LIST_EMPTY)
		}
		panic(err)
	}
	return reviews
}
