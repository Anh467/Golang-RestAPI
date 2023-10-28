package storage

import (
	"common"
	"context"
	"entities"
	"fmt"
)

func (s *sqlserverStore) CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error) {
	userJWTModel := &entities.UserJWTModel{}
	if err := s.db.Table(entities.UserModelTable).Create(user).Error; err != nil {
		return userJWTModel, err
	}
	tokenString, err := common.CreateToken(user)
	if err != nil {
		return nil, err
	}
	// res
	fmt.Printf("user: " + user.Email + " " + user.FullName + " " + user.Role + " " + tokenString)
	userJWTModel.UserID = user.UserID
	userJWTModel.Email = user.Email
	userJWTModel.FullName = user.FullName
	userJWTModel.Role = user.Role
	userJWTModel.Token = tokenString
	return userJWTModel, nil
}
