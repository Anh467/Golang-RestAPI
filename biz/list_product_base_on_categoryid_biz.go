package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b createBiz) ListProductBaseOnCategoryIDBiz(ctx context.Context, categoryid, offsetNum int, limitNum int) []entities.ProductModel {
	if categoryid <= 0 {
		panic(common.CATEGORY_MUST_POSITIVE)
	}
	products := b.store.ListProductBaseOnCategoryIDStorage(ctx, categoryid, offsetNum, limitNum)
	return products
}
