package biz

import (
	"common"
	"context"
	"entities"
)

func (b *createBiz) GetCartBiz(ctx context.Context, userid, productid int) *entities.CartGet {
	// check
	if productid < 0 {
		panic(common.PRODUCT_CANT_NEGATIVE)
	}
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// deleting
	cart := b.store.GetCartStorage(ctx, userid, productid)
	// res
	return cart
}
