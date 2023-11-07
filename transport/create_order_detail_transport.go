package transport

import (
	"main/biz"
	"main/common"
	"main/entities"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create /:userid/:orderid (sẽ lấy danh sách của cart)
// CreateOrderDetailStorage(ctx context.Context, userid, orderid int, orderdetails []entities.OrderDetailCreate) []entities.OrderDetailGet
func CreateOrderDetailTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var orderDetails []entities.OrderDetailCreate
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get body
		if err := c.ShouldBindJSON(&orderDetails); err != nil {
			panic(err)
		}
		//dependency
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// do create
		orderDetailsTemp := business.CreateOrderDetailBiz(c, userid, orderid, orderDetails)
		//res
		c.JSON(http.StatusOK, gin.H{
			"orderdetails": orderDetailsTemp,
		})
	}
}
