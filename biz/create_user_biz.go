package biz

import (
	"context"
	"entities"
	"errors"
	"regexp"
)

func (biz *createBiz) CreateUserBiz(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error) {
	// check blank
	if user.FullName == "" {
		return nil, errors.New(entities.FULL_NAME_BLANK)
	}
	if user.Password == "" {
		return nil, errors.New(entities.PASS_WORD_BLANK)
	}
	if user.Email == "" {
		return nil, errors.New(entities.EMAILL_BLANK)
	}
	// check regex
	match, err := regexp.MatchString(entities.EMAIL_REGEX, user.Email)
	if err != nil || match == false {
		return nil, errors.New(entities.EMAILL_WRONG_REGEX)
	}
	// return userRespone
	userJWTModel, err := biz.store.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return userJWTModel, nil
}
