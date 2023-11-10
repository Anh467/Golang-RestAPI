package biz

import (
	"context"
	"main/common"
)

func (b *createBiz) DeleteReviewBiz(ctx context.Context, reviewid, userid int) {
	// check negative reviewid
	if reviewid < 0 {
		panic(common.REVIEW_ID_CAN_NOT_BE_NEGATIVE)
	}
	// check negative reviewid
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// creating
	b.store.DeleteReviewStorage(ctx, reviewid, userid)
}
