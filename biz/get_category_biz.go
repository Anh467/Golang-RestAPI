package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) GetCategoryBiz(ctx context.Context, categoryid int) entities.CategoryModel {
	// business rule
	if categoryid == 0 {
		panic(common.CATEGORY_ID_BLANK)
	}
	// getting
	category := biz.store.GetCategory(ctx, categoryid)
	return category
}
