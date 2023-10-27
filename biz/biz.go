package biz

import (
	"context"
	"entities"
)

type CreateStorage interface {
	ListUser(ctx context.Context) ([]entities.UserModel, error)
}

type createBiz struct {
	store CreateStorage
}

func NewCreateBiz(store CreateStorage) *createBiz {
	return &createBiz{store: store}
}
