package biz

import (
	"context"
	"main/common"
)

func (b *createBiz) DeleteCartBiz(ctx context.Context, userid, productid int) {
	// check
	if productid < 0 {
		panic(common.PRODUCT_CANT_NEGATIVE)
	}
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// deleting
	b.store.DeleteCartStorage(ctx, userid, productid)
}
