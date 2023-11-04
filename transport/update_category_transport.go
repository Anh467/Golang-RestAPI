package transport

import (
	"main/biz"
	"main/common"
	"main/entities"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCategoryTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare catgory
		var category entities.CategoryUpdate
		var categoryid int
		var err error
		// get data from request body
		if err := c.ShouldBindJSON(&category); err != nil {
			panic(err)
		}
		// get slug from url
		categoryid, err = strconv.Atoi(c.Param("categoryid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// updating
		categoryTemp := business.UpdateCategoryBiz(c, category, categoryid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"category": categoryTemp,
		})
	}
}
