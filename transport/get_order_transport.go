package transport

import (
	"biz"
	"common"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrderTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get flag
		flag := c.GetBool("flag")
		// get header
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.CreateStorage(store)
		//  geting
		order := business.GetOrder(c, orderid, flag)
		// res
		c.JSON(http.StatusOK, gin.H{
			"order": order,
		})
	}
}
