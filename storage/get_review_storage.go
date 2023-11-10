package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) GetReviewStorage(ctx context.Context, reviewid int) entities.ReviewGet {
	// declare review
	var review entities.ReviewGet
	// check existion of the review
	var count int64
	s.db.Where("ReviewID = ?", reviewid).First(&entities.ReviewModel{}).Count(&count)
	if count == 0 {
		panic(common.REVIEW_ID_NOT_EXIST)
	}
	// get the review
	if err := s.db.Preload("User").Where("ReviewID = ?", reviewid).First(&review).Error; err != nil {
		panic(err)
	}
	return review
}
