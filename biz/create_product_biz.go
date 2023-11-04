package biz

import (
	"context"
	"main/common"
	"main/entities"
)

func (biz *createBiz) CreateProductBiz(ctx context.Context, product *entities.ProductModel) *entities.ProductModel {
	// check blank
	if product.Name == "" {
		panic(common.NAME_BLANK)
	}
	if product.Description == "" {
		panic(common.DESCRIPTION_BLANK)
	}
	if product.Price == 0 {
		panic(common.PRICE_BLANK)
	}
	if product.CategoryID == 0 {
		panic(common.CATEGORY_ID_BLANK)
	}
	// create product
	productTemp := biz.store.CreateProduct(ctx, product)
	// res
	return productTemp
}
