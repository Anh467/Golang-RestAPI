package biz

import (
	"common"
	"context"
	"entities"
)

func (biz *createBiz) UpdateCategoryBiz(ctx context.Context, category entities.CategoryUpdate, categoryid int) *entities.CategoryModel {
	// business rule
	if categoryid == 0 {
		panic(common.CATEGORY_ID_BLANK)
	}
	if category.Name == "" {
		panic(common.CATEGORY_NAME_BLANK)
	}
	// updating
	categoryTemp := biz.store.UpdateCategory(ctx, category, categoryid)
	// res
	return categoryTemp
}
