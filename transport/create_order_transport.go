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

func CreateOrderTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var order entities.OrderCreate
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get request
		c.ShouldBindJSON(&order)
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.CreateStorage(store)
		//  creating
		orderTemp := business.CreateOrder(c, order, userid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"order": orderTemp,
		})
	}
}
