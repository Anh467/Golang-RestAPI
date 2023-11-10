package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) UpdateReviewStorage(ctx context.Context, review entities.ReviewUpdate, userid int) entities.ReviewGet {

	// declare review
	var count int64
	reviewTemp := &entities.ReviewGet{
		ReviewID: review.ReviewID,
		Rating:   review.Rating,
		Comment:  review.Comment,
	}

	// check existion of the review .Where("UserID = ?", userid)
	s.db.Where("ReviewID = ? and UserID = ?", reviewTemp.ReviewID, userid).First(&entities.ReviewModel{}).Count(&count)
	if count == 0 {
		panic(common.REVIEW_ID_NOT_EXIST)
	}
	// get the review
	if err := s.db.Preload("User").Where("ReviewID = ?", reviewTemp.ReviewID).First(&reviewTemp).Error; err != nil {
		panic(err)
	}
	return *reviewTemp
}
