package storage

import (
	"context"
	"main/entities"
)

func (s *sqlserverStore) GetUserByID(ctx context.Context, userid int) entities.UserGetModel {
	var user entities.UserGetModel
	// get user by Password, Email
	if err := s.db.Where("UserID = ?", userid).Table(entities.UserModelTable).First(&user).Error; err != nil {
		panic(err)
	}

	return user
}
