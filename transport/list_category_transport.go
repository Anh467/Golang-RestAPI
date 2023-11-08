package transport

import (
	"main/biz"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCategoryTransport(db *gorm.DB) func(c *gin.Context) {
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
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// getting
		categories := business.ListCategoryBiz(c, offsetNum, limitNum)
		// res
		c.JSON(http.StatusOK, gin.H{
			"length":     len(categories),
			"categories": categories,
		})
	}
}
