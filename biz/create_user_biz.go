package biz

import (
	"context"
	"errors"
	"main/common"
	"main/entities"
	"regexp"
)

func (biz *createBiz) CreateUserBiz(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error) {
	// check blank
	if user.FullName == "" {
		return nil, errors.New(common.FULL_NAME_BLANK)
	}
	if user.Password == "" {
		return nil, errors.New(common.PASS_WORD_BLANK)
	}
	if user.Email == "" {
		return nil, errors.New(common.EMAILL_BLANK)
	}
	if user.Address == "" {
		return nil, errors.New(common.ADDRESS_BLANK)
	}
	// check regex
	match, err := regexp.MatchString(common.EMAIL_REGEX, user.Email)
	if err != nil || match == false {
		return nil, errors.New(common.EMAILL_WRONG_REGEX)
	}
	// return userRespone
	userJWTModel, err := biz.store.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return userJWTModel, nil
}
