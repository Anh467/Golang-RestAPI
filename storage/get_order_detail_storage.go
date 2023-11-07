package storage

import (
	"context"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) getOrderDetailStorage(ctx context.Context, userid, orderid int) []entities.OrderDetailGet {
	var orderDetailsGet []entities.OrderDetailGet
	// check this userid own this orderid
	countRows := s.db.Where("UserID = ?", userid).Where("OrderID = ?", orderid).
		Table(entities.OrderDetailModelTable).Find(nil).RowsAffected
	if countRows == 0 {
		panic("This userid not own this orderid")
	}
	// get
	if err := s.db.Where("OrderID = ?", orderid).
		Preload(entities.ProductModelTable, func(db *gorm.DB) *gorm.DB {
			return db.Select("ProductID", "Name", "Image")
		}).
		Find(&orderDetailsGet); err != nil {
		panic(err)
	}
	// response
	return orderDetailsGet
}
