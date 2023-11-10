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

func GetReviewStorage(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get headers
		reviewid, err := strconv.Atoi(c.Param("reviewid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// geting
		review := bussiness.GetReviewBiz(c, reviewid)
		c.JSON(http.StatusOK, gin.H{
			"review": review,
		})
	}
}
