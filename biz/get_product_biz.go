package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) GetProductBiz(ctx context.Context, productid string) *entities.ProductModel {
	if productid == "" {
		panic(common.PRODUCT_ID_BLANK)
	}
	product := biz.store.GetProduct(ctx, productid)

	return product
}
