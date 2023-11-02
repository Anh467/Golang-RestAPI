package biz

import (
	"common"
	"context"
	"entities"
)

func (b *createBiz) CreateOrder(ctx context.Context, order entities.OrderCreate, userid int) entities.OrderModel {
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
	// creating
	orderTemp := b.store.CreateOrder(ctx, order, userid)
	return orderTemp
}