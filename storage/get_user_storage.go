package storage

import (
	"context"
	"entities"
)

func (s *sqlserverStore) GetUser(ctx context.Context, email, pass string) entities.UserGetModel {
	var user entities.UserGetModel
	// get user by Password, Email
	if err := s.db.Where("Email", email).Where("Email", email).First(&user).Error; err != nil {
		panic(err)
	}

	return user
}
