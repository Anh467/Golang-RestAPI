package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) DeleteReviewStorage(ctx context.Context, reviewid, userid int) {
	// check existion of the review
	var count int64
	s.db.Where("ReviewID = ?", reviewid).Where("UserID = ?", userid).First(&entities.ReviewModel{}).Count(&count)
	if count == 0 {
		panic(common.REVIEW_ID_NOT_EXIST)
	}
	// delete the review
	tx := s.db.Where("ReviewID = ?", reviewid).Delete(&entities.ReviewModel{})
	if tx.RowsAffected == 0 {
		panic(common.REVIEW_MODIFY_FAIL)
	}
}
