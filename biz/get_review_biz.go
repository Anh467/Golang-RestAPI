package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) GetReviewBiz(ctx context.Context, reviewid int) entities.ReviewGet {
	// check negative
	if reviewid < 0 {
		panic(common.REVIEW_ID_CAN_NOT_BE_NEGATIVE)
	}
	// getting
	review := b.store.GetReviewStorage(ctx, reviewid)
	return review
}
