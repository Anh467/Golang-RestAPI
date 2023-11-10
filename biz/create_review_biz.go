package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) CreateReviewBiz(ctx context.Context, review entities.ReviewCreate) entities.ReviewGet {
	// check negative productid
	if review.ProductID < 0 {
		panic(common.PRODUCT_CANT_NEGATIVE)
	}
	// check negative userid
	if review.UserID < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// creating
	reviewTemp := b.store.CreateReviewStorage(ctx, review)
	return reviewTemp
}
