package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) UpdateReviewBiz(ctx context.Context, review entities.ReviewUpdate, userid int) entities.ReviewGet {
	// check negative reviewid
	if review.ReviewID < 0 {
		panic(common.REVIEW_ID_CAN_NOT_BE_NEGATIVE)
	}
	// creating
	reviewTemp := b.store.UpdateReviewStorage(ctx, review, userid)
	return reviewTemp
}
