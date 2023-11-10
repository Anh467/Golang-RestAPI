package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) CreateReviewStorage(ctx context.Context, review entities.ReviewCreate) entities.ReviewGet {
	// declate
	reviewGet := entities.ReviewGet{
		ProductID: review.ProductID,
		UserID:    review.UserID,
		Rating:    review.Rating,
		Comment:   review.Comment,
	}
	var count int64
	// check the existant of the userid
	s.db.Table(entities.UserModelTable).Where("UserID = ?", review.UserID).Count(&count)
	if count == 0 {
		panic(common.USER_ID_NOT_FOUND)
	}
	// check the existant of the productid
	s.db.Table(entities.ProductModelTable).Where("ProductID = ?", review.ProductID).Count(&count)
	if count == 0 {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
	// create new a review
	if err := s.db.Preload("User").Create(&reviewGet).Error; err != nil {
		panic(err)
	}
	return reviewGet
}
