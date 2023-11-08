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

func GetUserByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get header
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get role
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// get user
		user := business.GetUserByIDBiz(c, userid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
