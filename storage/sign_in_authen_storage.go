package storage

import (
	"common"
	"context"
	"entities"
)

func (s *sqlserverStore) SignInAuthen(ctx context.Context, email, pass string) (*entities.UserJWTModel, error) {
	userJWTModel := &entities.UserJWTModel{}
	user := &entities.UserCreateModel{}
	// get user
	if err := s.db.Table(entities.UserModelTable).Where("Email = ?", email).Where("Password = ?", pass).First(user).Error; err != nil {
		return nil, err
	}
	// create token
	tokenString, err := common.CreateToken(user)
	if err != nil {
		return nil, err
	}
	// res
	userJWTModel.UserID = user.UserID
	userJWTModel.Email = user.Email
	userJWTModel.FullName = user.FullName
	userJWTModel.Role = user.Role
	userJWTModel.Token = tokenString
	return userJWTModel, nil
}
