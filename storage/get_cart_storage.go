package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) GetCartStorage(ctx context.Context, userid, productid int) *entities.CartGet {
	// declare variables
	var cart *entities.CartGet
	// get the cart
	if err := s.db.Where("ProductID = ?", productid).Where("UserID = ?", userid).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CART_NOT_EXIST)
		}
		panic(err)
	}
	// get product information
	if err := s.db.Where("ProductID = ?", productid).First(&cart.Product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		}
		panic(err)
	}
	return cart
}
