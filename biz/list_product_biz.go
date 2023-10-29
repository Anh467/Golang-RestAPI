package biz

import (
	"context"
	"entities"
)

func (biz *createBiz) ListProductBiz(ctx context.Context, offsetNum int, limitNum int) ([]entities.Product, error) {
	//
	products, err := biz.store.ListProduct(ctx, offsetNum, limitNum)
	if err != nil {
		return nil, err
	}
	return products, nil
}
