package storage

import (
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) UpdateReviewStorage(ctx, review entities.ReviewUpdate) entities.ReviewGet {
	// declare review
	reviewTemp := &entities.ReviewGet{
		ReviewID: review.ReviewID,
		Rating:   review.Rating,
		Comment:  review.Comment,
	}
	// check existion of the review
	var count int64
	s.db.Where("ReviewID = ?", reviewTemp.ReviewID).First(&entities.ReviewModel{}).Count(&count)
	if count == 0 {
		panic(common.REVIEW_ID_NOT_EXIST)
	}
	// update the review
	if err := s.db.Where("ReviewID = ?", reviewTemp.ReviewID).Updates(&reviewTemp).Error; err != nil {
		panic(err)
	}
	// get the review
	if err := s.db.Preload("User").Where("ReviewID = ?", reviewTemp.ReviewID).First(&review).Error; err != nil {
		panic(err)
	}
	return *reviewTemp
}
