package biz

import (
	"common"
	"context"
	"entities"
)

func (b *createBiz) UpdateCartBiz(ctx context.Context, cart entities.CartUpdate) *entities.CartGet {
	// check blank
	if cart.Quantity < 0 {
		panic(common.CART_QUANTITY_CANT_NEGATIVE)
	}
	if cart.ProductID < 0 {
		panic(common.PRODUCT_CANT_NEGATIVE)
	}
	if cart.UserID < 0 {
		panic(common.USER_ID_CANT_NEGATIVE)
	}
	// creating
	cartTemp := b.store.UpdateCartStorage(ctx, cart)
	// res
	return cartTemp
}
