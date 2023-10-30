package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) GetProductBiz(ctx context.Context, productid string) *entities.Product {
	if productid == "" {
		panic(common.PRODUCT_ID_BLANK)
	}
	product := biz.store.GetProduct(ctx, productid)

	return product
}
