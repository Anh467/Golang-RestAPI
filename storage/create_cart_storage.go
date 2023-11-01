package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlserverStore) CreateCartStorage(ctx context.Context, cart entities.CartCreate) *entities.CartGet {
	// declare variables
	cartTemp := &entities.CartGet{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
	}

	// get the cart
	if err := s.db.Create(&cartTemp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.CART_NOT_EXIST)
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			panic(common.CART_USER_ID_AND_PRODUCT_ID_DUPILICATE)
		}
		panic(err)
	}
	// get product information
	if err := s.db.Where("ProductID = ?", cartTemp.ProductID).First(&cartTemp.Product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		}
		panic(err)
	}
	// res
	return cartTemp
}
