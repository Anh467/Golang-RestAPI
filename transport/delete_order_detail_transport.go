package transport

import (
	"main/biz"
	"main/common"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// delete/:userid/:orderid/:productid (Chỉ update khi trạng thái order là STATUS_WAIT_FOR_CONFIRMATION)
// DeleteOrderDetailStorage(ctx context.Context, userid, orderid, productid int, flag bool)
func DeleteOrderDetailTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
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
		//dependency
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// do delete
		business.DeleteOrderDetailBiz(c, userid, orderid, productid, flag)
		//res
		c.JSON(http.StatusOK, gin.H{
			"status": "done",
		})
	}
}
