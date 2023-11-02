package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) GetOrder(ctx context.Context, orderid int, flag bool) entities.OrderModel {
	var order entities.OrderModel
	// if flag is true that mean ur admin can able access to cancled orders
	if flag {
		if err := s.db.Where("OrderID = ?", orderid).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(common.ORDER_DELETE_FAIL)
			}
			panic(err)
		}
		return order
	}
	// this below mean this is just normal user and unable to access cancled orders
	if err := s.db.Where("OrderID = ?", orderid).
		Where("Status != ?", entities.STATUS_CANCELED).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.ORDER_DELETE_FAIL)
		}
		panic(err)
	}
	return order
}
