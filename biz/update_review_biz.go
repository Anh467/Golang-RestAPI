package biz

import (
	"main/common"
	"main/entities"
)

func (b *createBiz) UpdateReviewBiz(ctx, review entities.ReviewUpdate) entities.ReviewGet {
	// check negative reviewid
	if review.ReviewID < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// creating
	reviewTemp := b.store.UpdateReviewStorage(ctx, review)
	return reviewTemp
}
