package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateOrder(ctx context.Context, order entities.OrderCreate, userid int) entities.OrderModel {
	// declare
	orderTemp := entities.OrderModel{
		UserID: userid,
	}
	// this function below works simultaneously as creating, and checking the existant of user
	if err := s.db.Where("UserID = ?", orderTemp.UserID).First(&orderTemp.User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.USER_ID_NOT_FOUND)
		}
	}
	// creating order
	if err := s.db.Create(&orderTemp).Error; err != nil {
		panic(err)
	}
	// return the result after creating the order
	return orderTemp
}
