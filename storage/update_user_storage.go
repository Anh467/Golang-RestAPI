package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) UpdateUser(ctx context.Context, userUpdate entities.UserUpdate, userid int) *entities.UserModel {
	var userTemp *entities.UserModel
	// get user id
	if err := s.db.Where("UserID = ?", userid).Table(entities.UserModelTable).Updates(&userUpdate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.USER_ID_NOT_FOUND)
		} else {
			panic(err)
		}
	}
	// get updated data
	if err := s.db.Where("UserID = ?", userid).First(&userTemp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	return userTemp
}
