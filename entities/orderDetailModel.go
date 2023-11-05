package entities

type OrderDetailModel struct {
	OrderID   int          `json:"orderid" gorm:"column:OrderID"`
	ProductID int          `json:"productid" gorm:"column:ProductID"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity"`
	Order     OrderCreate  `json:"user" gorm:"foreignKey:OrderID"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
}

type OrderDetailCreate struct {
	OrderID   int `json:"orderid" gorm:"column:OrderID"`
	ProductID int `json:"productid" gorm:"column:ProductID"`
	Quantity  int `json:"quantity" gorm:"column:Quantity"`
}

type OrderDetailUpdate struct {
	Quantity int `json:"quantity" gorm:"column:Quantity"`
}
