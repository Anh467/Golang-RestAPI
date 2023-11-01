package biz

import (
	"common"
	"context"
	"entities"
)

func (b *createBiz) ListCartBiz(ctx context.Context, userid int) []entities.CartGet {
	// check
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// listing
	cartList := b.store.ListCartStorage(ctx, userid)
	// res
	return cartList
}
