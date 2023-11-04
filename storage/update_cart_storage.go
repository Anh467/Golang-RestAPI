package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) UpdateCartStorage(ctx context.Context, cart entities.CartUpdate) *entities.CartGet {
	// declare variables
	cartTemp := &entities.CartGet{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
	}

	// get the cart
	tx := s.db.Where("ProductID = ?", cartTemp.ProductID).Where("UserID = ?", cartTemp.UserID).Updates(&cartTemp)
	if tx.Error != nil {
		panic(tx.Error)
	}
	// check affected
	if tx.RowsAffected == 0 {
		panic(common.CART_NOT_EXIST)
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
