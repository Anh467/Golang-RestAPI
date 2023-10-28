package biz

import (
	"context"
	"entities"
)

type CreateStorage interface {
	ListUser(ctx context.Context) ([]entities.UserModel, error)
	CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error)
}

type createBiz struct {
	store CreateStorage
}

func NewCreateBiz(store CreateStorage) *createBiz {
	return &createBiz{store: store}
}
