package storage

import (
	"context"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateOrderDetailStorageg(ctx context.Context, userid, orderid int, orderdetails []entities.OrderDetailCreate) []entities.OrderDetailGet {
	// check this userid own this orderid
	countRows := s.db.Where("UserID = ?", userid).Where("OrderID = ?", orderid).
		Table(entities.OrderDetailModelTable).Find(nil).RowsAffected
	if countRows == 0 {
		panic("This userid not own this orderid")
	}
	// declare response
	var orderDetailsGet []entities.OrderDetailGet
	// init order details
	for index, _ := range orderdetails {
		orderdetails[index].OrderID = orderid
	}
	/// Create
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// rollback
		if err := tx.Create(&orderdetails).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
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
