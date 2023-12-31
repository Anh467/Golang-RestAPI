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

func DeleteOrderTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get header
		orderid, err := strconv.Atoi(c.Param("orderid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//  deleting
		business.DeleteOrderBiz(c, orderid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "done",
		})
	}
}
