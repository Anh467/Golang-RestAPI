package storage

import (
	"common"
	"context"
	"entities"
)

func (s *sqlserverStore) DeleteOrder(ctx context.Context, orderid int) {
	// deleting
	tx := s.db.Where("OrderID = ?", orderid).
		Table(entities.ORDER_MODEL_TABLE).
		Update(entities.ORDER_COLUMN_Status, entities.STATUS_CANCELED)
	// check errors
	if tx.Error != nil {
		panic(tx.Error)
	}
	// check deleted or not
	if tx.RowsAffected == 0 {
		panic(common.ORDER_DELETE_FAIL)
	}
}
