package biz

import (
	"common"
	"context"
	"entities"
)

func (b *createBiz) GetOrderBiz(ctx context.Context, orderid int, flag bool) entities.OrderModel {
	// check orderid
	if orderid <= 0 {
		panic(common.ORDER_ID_CANT_NEGATIVE)
	}
	// deleting
	order := b.store.GetOrder(ctx, orderid, flag)
	return order
}
