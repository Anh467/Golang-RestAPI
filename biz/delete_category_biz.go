package biz

import (
	"common"
	"context"
)

func (biz *createBiz) DeleteCategoryBiz(ctx context.Context, categoryid int) {
	// business rule
	if categoryid == 0 {
		panic(common.CATEGORY_ID_BLANK)
	}
	// deleting
	biz.store.DeleteCategory(ctx, categoryid)
}
