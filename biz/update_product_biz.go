package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) UpdateProductBiz(ctx context.Context, product entities.ProductUpdate, productid int) *entities.Product {
	// check blank
	if productid == 0 {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
	// update
	productTemp := biz.store.UpdateProduct(ctx, product, productid)
	return productTemp
}
