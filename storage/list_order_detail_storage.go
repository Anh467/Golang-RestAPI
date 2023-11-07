package storage

import (
	"context"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListOrderDetailStorage(ctx context.Context, userid, orderid, offset, limit int, flag bool) []entities.OrderDetailGet {
	// declare
	var orderdetails []entities.OrderDetailGet

	// check this userid admin or not
	if !flag {
		// check this userid own this orderid
		var countRows int64
		s.db.Where("UserID = ?", userid).Where("OrderID = ?", orderid).
			Table(entities.ORDER_MODEL_TABLE).Count(&countRows)
		if countRows == 0 {
			panic(common.ORDER_USER_NOT_OWN_THIS_ODER)
		}
	}
	// get
	if err := s.db.Where("OrderID = ?", orderid).
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("ProductID", "Name", "Image")
		}).
		Find(&orderdetails).Error; err != nil {
		panic(err)
	}
	// response
	return orderdetails
}
