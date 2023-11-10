package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListReviewBaseOnProductIDStorage(ctx context.Context, productid, offsetNum int, limitNum int) []entities.ReviewGet {
	// declare list
	var reviews []entities.ReviewGet
	// get list
	if err := s.db.Offset(offsetNum).Where("ProductID = ?", productid).Limit(limitNum).Preload("User").Find(&reviews).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.REVIEW_LIST_EMPTY)
		}
		panic(err)
	}
	return reviews
}
