package storage

import (
	"context"
	"entities"
)

func (s *sqlserverStore) CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error) {
	var userJWTModel *entities.UserJWTModel
	if err := s.db.Table(entities.CategorieModelTable).Create(user).Error; err != nil {
		return userJWTModel, err
	}
	userJWTModel.Email = user.Email
	userJWTModel.FullName = user.FullName
	userJWTModel.Role = user.Role
	userJWTModel.Token = ""
	return userJWTModel, nil
}
