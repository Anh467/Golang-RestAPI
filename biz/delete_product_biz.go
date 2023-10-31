package biz

import (
	"common"
	"context"
)

func (biz *createBiz) DeleteProductBiz(ctx context.Context, productid int) {
	// check blank
	if productid == 0 {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
	// update
	biz.store.DeleteProduct(ctx, productid)
}
