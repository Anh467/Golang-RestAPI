package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) CreateOrderDetailBiz(ctx context.Context, userid, orderid int, orderdetails []entities.OrderDetailCreate) []entities.OrderDetailGet {
	for _, orderdetail := range orderdetails {
		if orderdetail.Quantity <= 0 {
			panic(common.ORDER_DETAIL_QUANTITY_MUST_POSITIVE)
		}
	}
	orderdetailsget := biz.store.CreateOrderDetailStorage(ctx, userid, orderid, orderdetails)
	return orderdetailsget
}
