package storage

import (
	"context"
	"errors"
	"main/common"
	"main/entities"

	"gorm.io/gorm"
)

func (s *sqlserverStore) DeleteProduct(ctx context.Context, productid int) {
	tx := s.db.Where("ProductID = ?", productid).Table(entities.ProductModelTable).Delete(nil)
	// check error
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.PRODUCT_ID_NOT_EXIST)
		} else {
			panic(err)
		}
	}
	if tx.RowsAffected == 0 {
		panic(common.PRODUCT_ID_NOT_EXIST)
	}
}
