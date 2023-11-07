package biz

import (
	"context"
	"main/entities"
)

func (biz *createBiz) ListOrderDetail(ctx context.Context, userid, orderid, offset, limit int, flag bool) []entities.OrderDetailGet {
	orderdetailget := biz.store.ListOrderDetailStorage(ctx, userid, orderid, offset, limit, flag)
	return orderdetailget
}
