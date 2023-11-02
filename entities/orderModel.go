package entities

//table
const ORDER_MODEL_TABLE = "Orders"

// column
const ORDER_COLUMN_OrderID = "OrderID"
const ORDER_COLUMN_UserID = "UserID"
const ORDER_COLUMN_OrderDate = "OrderDate"
const ORDER_COLUMN_Status = "Status"

// status
const STATUS_DENIED string = "denied"          // denied is the term which mean that the user stop the delivering
const STATUS_DELIVERING string = "delivering"  // order is delivering to the user
const STATUS_WAIT_FOR_CONFIRMATION = "confirm" // after ordering user waits for confirmation of the admin
const STATUS_COMPLETED = "completed"           // order was completed delivering to the user
const STATUS_CANCELED = "canceled"             // User deleted the order and user can't see
// status list
var STATUS_LIST = []string{STATUS_DENIED, STATUS_DELIVERING, STATUS_WAIT_FOR_CONFIRMATION, STATUS_COMPLETED, STATUS_CANCELED}

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
