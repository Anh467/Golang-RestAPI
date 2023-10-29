package biz

import (
	"context"
	"entities"
)

func (biz *createBiz) ListNewUserModel(ctx context.Context) ([]entities.UserModel, error) {
	//
	data, err := biz.store.ListUser(ctx)
	if err != nil {
		return data, err
	}
	return data, nil
}
