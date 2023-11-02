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

func ListOrderTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// offset
		var offsetNum, limitNum int = 0, 5
		// get params from url
		limit := c.Query("limit")
		offset := c.Query("offset")
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
		// get flag
		flag := c.GetBool("flag")
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//  listing
		order := business.ListOrderBiz(c, userid, limitNum, offsetNum, flag)
		// res
		c.JSON(http.StatusOK, gin.H{
			"length": len(order),
			"orders": order,
		})
	}
}
