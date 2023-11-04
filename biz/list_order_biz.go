package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) ListOrderBiz(ctx context.Context, userid, limit, offset int, flag bool) []entities.OrderModel {
	// check userid
	if userid < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// get list
	orderList := b.store.ListOrder(ctx, userid, limit, offset, flag)
	return orderList
}
