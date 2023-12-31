package transport

import (
	"main/biz"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCartTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// offset
		var offsetNum, limitNum int = 0, 5
		// get userid
		userid := c.MustGet("userid").(int)
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
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// creating
		cartList := business.ListCartBiz(c, userid, limitNum, offsetNum)
		// res
		c.JSON(http.StatusOK, gin.H{
			"cart":   cartList,
			"length": len(cartList),
		})
	}
}
