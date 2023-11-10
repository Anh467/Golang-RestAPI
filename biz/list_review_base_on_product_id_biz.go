package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) ListReviewBaseOnProductIDBiz(ctx context.Context, productid, offsetNum int, limitNum int) []entities.ReviewGet {
	// check negative productid
	if productid < 0 {
		panic(common.PRODUCT_CANT_NEGATIVE)
	}
	review := b.store.ListReviewBaseOnProductIDStorage(ctx, productid, offsetNum, limitNum)
	return review
}
