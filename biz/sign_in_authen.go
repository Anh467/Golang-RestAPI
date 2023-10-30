package biz

import (
	"common"
	"context"
	"entities"
	"errors"
)

func (biz *createBiz) SignInAuthen(ctx context.Context, email, pass string) (*entities.UserJWTModel, error) {
	// check blank
	if pass == "" {
		return nil, errors.New(common.PASS_WORD_BLANK)
	}
	if email == "" {
		return nil, errors.New(common.EMAILL_BLANK)
	}
	// return userRespone
	userJWTModel, err := biz.store.SignInAuthen(ctx, email, pass)
	if err != nil {
		return nil, err
	}
	return userJWTModel, nil
}
