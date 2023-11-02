package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateOrder(ctx context.Context, order entities.OrderCreate, userid int) entities.OrderModel {
	// set user
	order.UserID = userid
	order.OrderID = 0
	// declare
	orderTemp := entities.OrderModel{
		UserID: userid,
	}

	// creating order
	if err := s.db.Create(&order).Error; err != nil {
		panic(err)
	}
	// getting order
	if err := s.db.Where("UserID = ?", userid).Where("OrderID = ?", order.OrderID).First(&orderTemp).Error; err != nil {
		panic(err)
	}
	// this function below works simultaneously as creating, and checking the existant of user
	if err := s.db.Where("UserID = ?", userid).First(&orderTemp.User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.USER_ID_NOT_FOUND)
		}
	}
	// return the result after creating the order
	return orderTemp
}
