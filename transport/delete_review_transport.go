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

func DeleteReviewTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get headers
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		reviewid, err := strconv.Atoi(c.Param("reviewid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// delete
		bussiness.DeleteReviewBiz(c, reviewid, userid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "done",
		})
	}
}
