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

// List /:userid/:orderid (lấy danh sách product trong hóa đơn đó)
// ListOrderDetailStorage(ctx context.Context, userid, orderid, offset, limit int, flag bool) []entities.OrderDetailGet
func ListOrderDetailTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// offset limit default
		var offsetNum, limitNum int = 0, 5
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get params from url
		limit := c.Query("limit")
		offset := c.Query("offset")
		// get flag
		flag := c.GetBool("flag")
		// convert to string
		if offset != "" {
			i, err := strconv.Atoi(offset)
			if err == nil {
				offsetNum = i
			}
		}
		if limit != "" {
			i, err := strconv.Atoi(limit)
			if err == nil {
				limitNum = i
			}
		}
		//dependency
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// listing
		orderDetails := business.ListOrderDetailBiz(c, userid, orderid, offsetNum, limitNum, flag)
		//res
		c.JSON(http.StatusOK, gin.H{
			"orderdetails": orderDetails,
		})
	}
}
