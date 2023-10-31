package storage

import (
	"common"
	"context"
	"entities"
	"errors"

	"gorm.io/gorm"
)

func (store *sqlserverStore) DeleteUser(ctx context.Context, userid int) {
	var user *entities.UserGetModel
	// get user
	if err := store.db.Select("UserID", "Role").Where("UserID = ?", userid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(common.USER_ID_NOT_FOUND)
		} else {
			panic(err)
		}
	}
	// check this is deleted or not
	if user.Role == entities.ROLE_DELETED {
		panic(common.USER_ID_NOT_FOUND)
	}
	// action delete
	if err := store.db.Where("UserID = ?", user.UserID).Table(entities.UserModelTable).Update("Role", entities.ROLE_DELETED).Error; err != nil {
		panic(err)
	}
}
