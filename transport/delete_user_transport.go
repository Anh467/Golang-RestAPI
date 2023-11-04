package transport

import (
	"main/biz"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get data param url
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(err)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// delete user
		bussiness.DeleteUserBiz(c, userid)
		c.JSON(http.StatusOK, gin.H{
			"status": "delete done",
		})
	}
}
