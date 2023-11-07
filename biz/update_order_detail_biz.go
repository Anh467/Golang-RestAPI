package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) UpdateOrderDetailBiz(ctx context.Context, orderdetail entities.OrderDetailUpdate, userid, orderid, productid int, flag bool) entities.OrderDetailGet {
	if orderdetail.Quantity <= 0 {
		panic(common.ORDER_DETAIL_QUANTITY_MUST_POSITIVE)
	}
	orderdetailget := biz.store.UpdateOrderDetailStorage(ctx, orderdetail, userid, orderid, productid, flag)
	return orderdetailget
}
