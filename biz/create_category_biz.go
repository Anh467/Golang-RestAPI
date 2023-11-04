package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) CreateCategoryBiz(ctx context.Context, category entities.CategoryCreate) *entities.CategoryModel {
	// business rule
	if category.Name == "" {
		panic(common.CATEGORY_NAME_BLANK)
	}

	// creating
	categoryTemp := biz.store.CreateCategory(ctx, category)

	// res
	return categoryTemp
}
