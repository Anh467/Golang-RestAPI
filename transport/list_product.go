package transport

import (
	"biz"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProduct(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// default offset and limit
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
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//
		products, err := business.ListProductBiz(c, offsetNum, limitNum)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	}
}
