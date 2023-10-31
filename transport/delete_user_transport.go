package transport

import (
	"biz"
	"net/http"
	"storage"
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
		bussiness := biz.CreateStorage(store)
		// delete user
		bussiness.DeleteUser(c, userid)
		c.JSON(http.StatusOK, gin.H{
			"status": "delete done",
		})
	}
}
