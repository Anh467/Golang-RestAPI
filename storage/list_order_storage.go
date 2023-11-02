package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListOrder(ctx context.Context, userid, limit, offset int, flag bool) []entities.OrderModel {
	var order []entities.OrderModel
	// if flag is true that mean ur admin can able access to cancled orders
	if flag {
		if userid == 0 {
			if err := s.db.Limit(limit).Offset(limit * offset).Find(&order).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					panic(common.ORDER_DELETE_FAIL)
				}
				panic(err)
			}
		} else {
			if err := s.db.Limit(limit).Offset(limit*offset).Where("Userid = ?", userid).Find(&order).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					panic(common.ORDER_DELETE_FAIL)
				}
				panic(err)
			}
		}
		return order
	}
	// this below mean this is just normal user and unable to access cancled orders
	if err := s.db.Where("Userid = ?", userid).
		Where("Status != ?", entities.STATUS_CANCELED).
		Limit(limit).Offset(limit * offset).
		Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.ORDER_DELETE_FAIL)
		}
		panic(err)
	}
	return order
}
