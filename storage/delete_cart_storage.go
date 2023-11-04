package storage

import (
	"context"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) DeleteCartStorage(ctx context.Context, userid, productid int) {
	// delete
	tx := s.db.Where("UserID = ?", userid).Where("ProductID = ?", productid).Table(entities.CART_MODEL_TABLE).Delete(nil)
	if tx.Error != nil {
		panic(tx.Error)
	}
	// check affected
	if tx.RowsAffected == 0 {
		panic(common.CART_NOT_EXIST)
	}
}
