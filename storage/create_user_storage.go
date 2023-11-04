package storage

import (
	"context"
	"errors"
	"fmt"
	"main/common"
	"main/entities"
)

func (s *sqlserverStore) CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error) {
	userJWTModel := &entities.UserJWTModel{}
	// check unqiue email
	tx := s.db.Select("Email").Table(entities.UserModelTable).Limit(1).Where("Email = ?", user.Email).Find(userJWTModel)
	if tx.RowsAffected > 0 {
		return nil, errors.New(common.EMAILL_DUPLICATE)
	}
	// create user
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
