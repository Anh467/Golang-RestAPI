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

func ListReviewBaseOnProductIDTransport(db *gorm.DB) func(c *gin.Context) {
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
		// get headers
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// update
		reviews := bussiness.ListReviewBaseOnProductIDBiz(c, productid, offsetNum, limitNum)
		// res
		c.JSON(http.StatusOK, gin.H{
			"reviews": reviews,
		})
	}
}
