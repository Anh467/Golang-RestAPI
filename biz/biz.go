package biz

import (
	"context"
	"entities"
)

type CreateStorage interface {
	ListUser(ctx context.Context) ([]entities.UserModel, error)
	CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error)
	ListProduct(ctx context.Context, offsetNum int, limitNum int) ([]entities.Product, error)
}

type createBiz struct {
	store CreateStorage
}

func NewCreateBiz(store CreateStorage) *createBiz {
	return &createBiz{store: store}
}
