package storage

import (
	"context"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateOrderDetailStorage(ctx context.Context, userid, orderid int, orderdetails []entities.OrderDetailCreate) []entities.OrderDetailGet {
	// check this userid own this orderid
	var countRows int64
	if err := s.db.Table(entities.ORDER_MODEL_TABLE).Where("UserID = ? AND OrderID = ?", userid, orderid).Count(&countRows).Error; err != nil {
		panic(err)
	}
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
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("ProductID", "Name", "Image")
		}).
		Find(&orderDetailsGet).Error; err != nil {
		panic(err)
	}
	// response
	return orderDetailsGet
}
