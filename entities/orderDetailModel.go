package entities

// column
const OrderDetail_COLUMN_OrderID = "OrderID"
const OrderDetail_COLUMN_ProductID = "ProductID"
const OrderDetail_COLUMN_Quantity = "Quantity"

// table
const OrderDetailModelTable = "Products"

type OrderDetailModel struct {
	OrderID   int          `json:"orderid" gorm:"column:OrderID"`
	ProductID int          `json:"productid" gorm:"column:ProductID"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity"`
	Order     OrderCreate  `json:"user" gorm:"foreignKey:OrderID"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
}

type OrderDetailGet struct {
	OrderID   int          `json:"orderid" gorm:"column:OrderID"`
	ProductID int          `json:"productid" gorm:"column:ProductID"`
	Quantity  int          `json:"quantity" gorm:"column:Quantity"`
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

func (OrderDetailModel) TableName() string {
	return OrderDetailModelTable
}

func (OrderDetailCreate) TableName() string {
	return OrderDetailModelTable
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetailModelTable
}

func (OrderDetailGet) TableName() string {
	return OrderDetailModelTable
}
