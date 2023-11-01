package entities

const CART_MODEL_TABLE = "Cart"

type CartModel struct {
	UserID    int          `json:"userid" gorm:"column:UserID;"`
	ProductID int          `json:"productid" gorm:"column:ProductID;"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity;"`
	User      UserGetModel `json:"user" gorm:"foreignKey:UserID"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
}

type CartCreate struct {
	Quantity int `json:"quantity" gorm:"column:Quantity;"`
}

type CartUpdate struct {
	Quantity int `json:"quantity" gorm:"column:Quantity;"`
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
