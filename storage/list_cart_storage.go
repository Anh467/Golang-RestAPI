package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) ListCartStorage(ctx context.Context, userid int) []entities.CartGet {
	// declare variables
	var cartList []entities.CartGet
	// get the cart
	if err := s.db.Where("UserID = ?", userid).First(&cartList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CART_NOT_EXIST)
		}
		panic(err)
	}
	// get product list information
	for index, ele := range cartList {
		productid := ele.ProductID
		if err := s.db.Where("ProductID = ?", productid).First(&cartList[index].Product).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(common.PRODUCT_ID_NOT_EXIST)
			}
			panic(err)
		}
	}
	// res
	return cartList
}
