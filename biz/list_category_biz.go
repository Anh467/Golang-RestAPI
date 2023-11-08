package biz

import (
	"context"
	"main/entities"
)

func (b *createBiz) ListCategoryBiz(ctx context.Context, offset, limit int) []entities.CategoryModel {
	Categories := b.store.ListCategoryStorage(ctx, offset, limit)
	return Categories
}
