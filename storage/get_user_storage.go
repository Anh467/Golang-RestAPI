package storage

import (
	"context"
	"main/entities"
)

func (s *sqlserverStore) GetUser(ctx context.Context, email, pass string) entities.UserGetModel {
	var user entities.UserGetModel
	// get user by Password, Email
	if err := s.db.Where("Email = ?", email).Where("Password = ?", pass).Table(entities.UserModelTable).First(&user).Error; err != nil {
		panic(err)
	}

	return user
}
