package biz

import (
	"context"
	"entities"
)

type CreateStorage interface {
	// user
	ListUser(ctx context.Context) ([]entities.UserModel, error)
	GetUser(ctx context.Context, email, pass string) entities.UserGetModel
	// authen
	CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error)
	SignInAuthen(ctx context.Context, email, pass string) (*entities.UserJWTModel, error)
	// product
	ListProduct(ctx context.Context, offsetNum int, limitNum int) ([]entities.Product, error)
	GetProduct(ctx context.Context, productid string) *entities.Product
	CreateProduct(ctx context.Context, product *entities.Product) *entities.Product
	DeleteProduct(ctx context.Context, productid int)
	UpdateProduct(ctx context.Context, product entities.ProductUpdate, productid int) *entities.Product
}

type createBiz struct {
	store CreateStorage
}

func NewCreateBiz(store CreateStorage) *createBiz {
	return &createBiz{store: store}
}
