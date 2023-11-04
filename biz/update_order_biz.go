package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (b *createBiz) UpdateOrderBiz(ctx context.Context, order entities.OrderUpdate, userid, orderid int) entities.OrderModel {
	// declare flag
	flag := false
	// check status is exist
	for _, ele := range entities.STATUS_LIST {
		if ele == order.Status {
			flag = true
			break
		}
	}
	// check status is exist or not
	if !flag {
		panic(common.STATUS_WRONG)
	}
	// check userid
	if userid <= 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// check orderid
	if orderid <= 0 {
		panic(common.ORDER_ID_CANT_NEGATIVE)
	}
	// updating
	orderTemp := b.store.UpdateOrder(ctx, order, userid, orderid)
	return orderTemp
}
