package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s sqlserverStore) DeleteOrderDetailStorage(ctx context.Context, userid, orderid, productid int, flag bool) {
	// get order status
	var order entities.OrderModel
	if err := s.db.Select("Status").Where("OrderID = ?", orderid).First(&order).Error; err != nil {
		panic(err)
	}
	// check this userid admin or not
	if !flag {
		// STATUS_WAIT_FOR_CONFIRMATION
		if order.Status != entities.STATUS_WAIT_FOR_CONFIRMATION {
			panic(common.ORDER_DETAIL_CANT_UPDATE)
		}
		// check this userid own this orderid
		var countRows int64
		s.db.Where("UserID = ?", userid).Where("OrderID = ?", orderid).
			Table(entities.ORDER_MODEL_TABLE).Count(&countRows)
		if countRows == 0 {
			panic("This userid not own this orderid")
		}
	}

	// delete
	tx := s.db.Where("OrderID = ?", orderid).Where("ProductID = ?", productid).Delete(&entities.OrderDetailModel{})
	// check err
	if tx.Error != nil {
		panic(tx.Error)
	}
	// check deleted or there is no order deleted?
	if tx.RowsAffected == 0 {
		panic(common.ORDER_DETAIL_NO_ROW_DELETE)
	}
}
