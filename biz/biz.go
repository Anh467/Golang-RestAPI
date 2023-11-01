package biz

import (
	"context"
	"entities"
)

type CreateStorage interface {
	// users
	ListUser(ctx context.Context) ([]entities.UserModel, error)
	GetUser(ctx context.Context, email, pass string) entities.UserGetModel
	UpdateUser(ctx context.Context, userUpdate entities.UserUpdate, userid int) *entities.UserModel
	DeleteUser(ctx context.Context, userid int)
	// authens
	CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error)
	SignInAuthen(ctx context.Context, email, pass string) (*entities.UserJWTModel, error)
	// producta
	ListProduct(ctx context.Context, offsetNum int, limitNum int) ([]entities.Product, error)
	GetProduct(ctx context.Context, productid string) *entities.Product
	CreateProduct(ctx context.Context, product *entities.Product) *entities.Product
	DeleteProduct(ctx context.Context, productid int)
	UpdateProduct(ctx context.Context, product entities.ProductUpdate, productid int) *entities.Product
	// categories
	UpdateCategory(ctx context.Context, category entities.CategoryUpdate, categoryid int) *entities.CategoryModel
	GetCategory(ctx context.Context, categoryid int) entities.CategoryModel
	DeleteCategory(ctx context.Context, categoryid int)
	CreateCategory(ctx context.Context, category entities.CategoryCreate) *entities.CategoryModel
}

type createBiz struct {
	store CreateStorage
}

func NewCreateBiz(store CreateStorage) *createBiz {
	return &createBiz{store: store}
}
