package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) ListCartBiz(ctx context.Context, userid, limit, offset int) []entities.CartGet {
	// check
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	if limit < 0 {
		panic(common.LIMIT_NOT_NEGATIVE)
	}
	if offset < 0 {
		panic(common.OFFSET_NOT_NEGATIVE)
	}
	// listing
	cartList := b.store.ListCartStorage(ctx, userid, limit, offset)
	// res
	return cartList
}
