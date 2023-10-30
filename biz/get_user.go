package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) GetUserBiz(ctx context.Context, email, pass string) entities.UserGetModel {
	var user entities.UserGetModel
	if email != "" {
		panic(common.EMAILL_BLANK)
	}

	if pass != "" {
		panic(common.PASS_WORD_BLANK)
	}
	return user
}
