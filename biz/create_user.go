package biz

import (
	"context"
	"entities"
	"errors"
)

type CreateUserModelStorage interface {
	CreateItem(ctx context.Context, data *entities.UserModel) error
}

func NewCreateToDoItemBiz(store CreateUserModelStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewUserModel(ctx context.Context, data *entities.UserModel) error {
	if data.FullName == "" {
		return errors.New("Name of an user can not be blank")
	}

	if data.Password == "" {
		return errors.New("Pass of an user can not be blank")
	}

	if data.Password == "" {
		return errors.New("Pass of an user can not be blank")
	}
	//
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
