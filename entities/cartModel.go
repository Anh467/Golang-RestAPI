package entities

//table
const CART_MODEL_TABLE = "Cart"

// column
const CART_COLUMN_QUANTITY = "Quantity"
const CART_COLUMN_PRODUCT_ID = "ProductID"

type CartModel struct {
	UserID    int          `json:"userid" gorm:"column:UserID;"`
	ProductID int          `json:"productid" gorm:"column:ProductID;"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity;"`
	User      UserGetModel `json:"user" gorm:"foreignKey:UserID"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
}
type CartGet struct {
	UserID    int          `json:"userid" gorm:"column:UserID;"`
	ProductID int          `json:"productid" gorm:"column:ProductID;"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity;"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
}
type CartBody struct {
	Quantity int `json:"quantity" gorm:"column:Quantity;"`
}

type CartCreate struct {
	UserID    int `json:"userid" gorm:"column:UserID;"`
	ProductID int `json:"productid" gorm:"column:ProductID;"`
	Quantity  int `json:"quantity" gorm:"column:Quantity;"`
}

type CartUpdate struct {
	UserID    int `json:"userid" gorm:"column:UserID;"`
	ProductID int `json:"productid" gorm:"column:ProductID;"`
	Quantity  int `json:"quantity" gorm:"column:Quantity;"`
}

func (CartModel) TableName() string {
	return CART_MODEL_TABLE
}

func (CartCreate) TableName() string {
	return CART_MODEL_TABLE
}
func (CartUpdate) TableName() string {
	return CART_MODEL_TABLE
}
func (CartGet) TableName() string {
	return CART_MODEL_TABLE
}

func (CartBody) TableName() string {
	return CART_MODEL_TABLE
}
