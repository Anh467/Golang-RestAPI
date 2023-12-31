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

func DeleteCategoryTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare catgory
		var categoryid int
		var err error
		// get slug from url
		categoryid, err = strconv.Atoi(c.Param("categoryid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// deleting
		business.DeleteCategoryBiz(c, categoryid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "done",
		})
	}
}
