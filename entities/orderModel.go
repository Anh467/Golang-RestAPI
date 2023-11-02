package entities

//table
const ORDER_MODEL_TABLE = "Orders"

// column
const ORDER_COLUMN_OrderID = "OrderID"
const ORDER_COLUMN_UserID = "UserID"
const ORDER_COLUMN_OrderDate = "OrderDate"
const ORDER_COLUMN_Status = "Status"

// status
const STATUS_DENIED = "denied"
const STATUS_DELIVERING = "delivering"
const STATUS_WAIT_FOR_CONFIRMATION = "confirm"
const STATUS_COMPLETED = "completed"

type OrderModel struct {
	OrderID   int          `json:"orderid" gorm:"column:OrderID;primaryKey"`
	UserID    int          `json:"userid" gorm:"column:UserID"`
	OrderDate string       `json:"orderdate" gorm:"column:OrderDate; default:GETDATE()"`
	Status    string       `json:"status" gorm:"column:Status; default:confirm"`
	User      UserGetModel `json:"user" gorm:"foreignKey:UserID"`
}

type OrderUpdate struct {
	Status string `json:"status" gorm:"column:Status"`
}

type OrderCreate struct {
	OrderID   int    `json:"orderid" gorm:"column:OrderID;primaryKey"`
	UserID    int    `json:"userid" gorm:"column:UserID"`
	OrderDate string `json:"orderdate" gorm:"column:OrderDate"`
	Status    string `json:"status" gorm:"column:Status; default:confirm"`
}

func (OrderModel) TableName() string {
	return ORDER_MODEL_TABLE
}
func (OrderUpdate) TableName() string {
	return ORDER_MODEL_TABLE
}

func (OrderCreate) TableName() string {
	return ORDER_MODEL_TABLE
}
