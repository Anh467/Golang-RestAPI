package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) UpdateProductBiz(ctx context.Context, product entities.ProductUpdate, productid int) *entities.ProductModel {
	// check blank
	if productid == 0 {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
	// update
	productTemp := biz.store.UpdateProduct(ctx, product, productid)
	return productTemp
}
