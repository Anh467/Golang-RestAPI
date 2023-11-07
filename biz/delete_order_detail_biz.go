package biz

import (
	"context"
)

func (biz *createBiz) DeleteOrderDetailBiz(ctx context.Context, userid, orderid, productid int, flag bool) {
	biz.store.DeleteOrderDetailStorage(ctx, userid, orderid, productid, flag)
}
