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

// Update /:userid/:orderid/:productid (Chỉ update khi trạng thái order là STATUS_WAIT_FOR_CONFIRMATION)
// UpdateOrderDetailStorage(ctx context.Context, orderdetail entities.OrderDetailUpdate, userid, orderid, productid int, flag bool) entities.OrderDetailGet
func UpdateOrderDetailTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var orderdetail entities.OrderDetailUpdate
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get flag
		flag := c.GetBool("flag")
		// get body
		if err := c.ShouldBindJSON(&orderdetail); err != nil {
			panic(common.JSON_BODY_WRONG_FORMAMT)
		}
		//dependency
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// do create
		orderDetailTemp := business.UpdateOrderDetailBiz(c, orderdetail, userid, orderid, productid, flag)
		//res
		c.JSON(http.StatusOK, gin.H{
			"orderdetail": orderDetailTemp,
		})
	}
}
