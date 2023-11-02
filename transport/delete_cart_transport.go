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

func DeleteCartTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get param from header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// creating
		business.DeleteCartBiz(c, userid, productid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "done",
		})
	}
}
