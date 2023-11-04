package biz

import (
	"context"
	"main/common"
)

func (biz *createBiz) DeleteUserBiz(ctx context.Context, userid int) {
	// check blank
	if userid == 0 {
		panic(common.USER_ID_BLANK)
	}
	// update
	biz.store.DeleteUser(ctx, userid)
}
