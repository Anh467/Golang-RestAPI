package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) GetUserBiz(ctx context.Context, email, pass string) entities.UserGetModel {
	if email == "" {
		panic(common.EMAILL_BLANK)
	}

	if pass == "" {
		panic(common.PASS_WORD_BLANK)
	}
	user := biz.store.GetUser(ctx, email, pass)

	return user
}
