package biz

import (
	"context"
	"main/entities"
)

func (b *createBiz) ListReviewAllBiz(ctx context.Context, offsetNum int, limitNum int) []entities.ReviewGet {
	review := b.store.ListReviewAllStorage(ctx, offsetNum, limitNum)
	return review
}
