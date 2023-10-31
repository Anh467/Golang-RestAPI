package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) UpdateUserStorage(ctx context.Context, userUpdate entities.UserUpdate, userid int) *entities.UserModel {
	// check blank
	if userid == 0 {
		panic(common.USER_ID_BLANK)
	}
	// update
	user := biz.store.UpdateUser(ctx, userUpdate, userid)
	return user
}
