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

func ListProductBaseOnCategoryIDTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		// default offset and limit
		var offsetNum, limitNum int = 0, 5
		// get slug from url
		categoryid, err := strconv.Atoi(c.Param("categoryid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
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
		products := business.ListProductBaseOnCategoryIDBiz(c, categoryid, offsetNum, limitNum)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			panic(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"length":   len(products),
			"products": products,
		})
	}
}
