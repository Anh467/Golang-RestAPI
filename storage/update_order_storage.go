package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) UpdateOrder(ctx context.Context, order entities.OrderUpdate, userid, orderid int) entities.OrderModel {
	orderTemp := entities.OrderModel{
		Status: order.Status,
	}
	// this function below works simultaneously as creating, and checking the existant of user
	if err := s.db.Where("UserID = ?", userid).First(&orderTemp.User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.USER_ID_NOT_FOUND)
		}
	}
	// updating
	if err := s.db.Where("UserID = ?", userid).
		Where("OrderID = ?", orderid).
		Updates(&order).
		Error; err != nil {
		panic(err)
	}
	// getting
	if err := s.db.Where("UserID = ?", userid).
		Where("OrderID = ?", orderid).
		First(&orderTemp).
		Error; err != nil {
		panic(err)
	}
	// res
	return orderTemp
}
