package transport

import (
	"biz"
	"common"
	"entities"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateOrderTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var order entities.OrderUpdate
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get request
		c.ShouldBindJSON(&order)
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//  updating
		orderTemp := business.UpdateOrderBiz(c, order, userid, orderid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"order": orderTemp,
		})
	}
}
