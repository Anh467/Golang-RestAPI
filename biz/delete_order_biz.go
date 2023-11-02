package biz

import (
	"common"
	"context"
)

func (b *createBiz) DeleteOrder(ctx context.Context, orderid int) {
	// check orderid
	if orderid <= 0 {
		panic(common.ORDER_ID_CANT_NEGATIVE)
	}
	// deleting
	b.store.DeleteOrder(ctx, orderid)
}
