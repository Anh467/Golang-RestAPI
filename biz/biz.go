package biz

import (
	"context"
	"main/entities"
)

type createStorage interface {
	// users
	ListUser(ctx context.Context) ([]entities.UserModel, error)
	GetUser(ctx context.Context, email, pass string) entities.UserGetModel
	UpdateUser(ctx context.Context, userUpdate entities.UserUpdate, userid int) *entities.UserModel
	DeleteUser(ctx context.Context, userid int)
	GetUserByID(ctx context.Context, userid int) entities.UserGetModel
	// authens
	CreateUser(ctx context.Context, user *entities.UserCreateModel) (*entities.UserJWTModel, error)
	SignInAuthen(ctx context.Context, email, pass string) (*entities.UserJWTModel, error)
	// producta
	ListProduct(ctx context.Context, offsetNum int, limitNum int) ([]entities.ProductModel, error)
	GetProduct(ctx context.Context, productid string) *entities.ProductModel
	CreateProduct(ctx context.Context, product *entities.ProductModel) *entities.ProductModel
	DeleteProduct(ctx context.Context, productid int)
	UpdateProduct(ctx context.Context, product entities.ProductUpdate, productid int) *entities.ProductModel
	ListProductBaseOnCategoryIDStorage(ctx context.Context, categoryid, offsetNum int, limitNum int) []entities.ProductModel
	// categories
	UpdateCategory(ctx context.Context, category entities.CategoryUpdate, categoryid int) *entities.CategoryModel
	GetCategory(ctx context.Context, categoryid int) entities.CategoryModel
	DeleteCategory(ctx context.Context, categoryid int)
	CreateCategory(ctx context.Context, category entities.CategoryCreate) *entities.CategoryModel
	ListCategoryStorage(ctx context.Context, offset, limit int) []entities.CategoryModel
	// cart
	CreateCartStorage(ctx context.Context, cart entities.CartCreate) *entities.CartGet
	DeleteCartStorage(ctx context.Context, userid, productid int)
	GetCartStorage(ctx context.Context, userid, productid int) *entities.CartGet
	ListCartStorage(ctx context.Context, userid, limit, offset int) []entities.CartGet
	UpdateCartStorage(ctx context.Context, cart entities.CartUpdate) *entities.CartGet
	// orders
	CreateOrder(ctx context.Context, order entities.OrderCreate, userid int) entities.OrderModel
	UpdateOrder(ctx context.Context, order entities.OrderUpdate, userid, orderid int) entities.OrderModel
	DeleteOrder(ctx context.Context, orderid int)
	GetOrder(ctx context.Context, orderid int, flag bool) entities.OrderModel
	ListOrder(ctx context.Context, userid, limit, offset int, flag bool) []entities.OrderModel
	// order details
	CreateOrderDetailStorage(ctx context.Context, userid, orderid int, orderdetails []entities.OrderDetailCreate) []entities.OrderDetailGet
	UpdateOrderDetailStorage(ctx context.Context, orderdetail entities.OrderDetailUpdate, userid, orderid, productid int, flag bool) entities.OrderDetailGet
	DeleteOrderDetailStorage(ctx context.Context, userid, orderid, productid int, flag bool)
	ListOrderDetailStorage(ctx context.Context, userid, orderid, offset, limit int, flag bool) []entities.OrderDetailGet
	// reviews
	CreateReviewStorage(ctx context.Context, review entities.ReviewCreate) entities.ReviewGet
	DeleteReviewStorage(ctx context.Context, reviewid, userid int)
	GetReviewStorage(ctx context.Context, reviewid int) entities.ReviewGet
	ListReviewAllStorage(ctx context.Context, offsetNum int, limitNum int) []entities.ReviewGet
	ListReviewBaseOnProductIDStorage(ctx context.Context, productid, offsetNum int, limitNum int) []entities.ReviewGet
	UpdateReviewStorage(ctx context.Context, review entities.ReviewUpdate, userid int) entities.ReviewGet
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
