package biz

import (
	"context"
	"main/entities"
)

func (biz *createBiz) GetUserByIDBiz(ctx context.Context, userid int) entities.UserGetModel {
	user := biz.store.GetUserByID(ctx, userid)

	return user
}
