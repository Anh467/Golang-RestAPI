package transport

import (
	"biz"
	"entities"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user entities.UserUpdate
		// get data param url
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(err)
		}
		// get request body data
		c.ShouldBindJSON(&user)
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// delete user
		userTemp := bussiness.UpdateUserBiz(c, user, userid)
		c.JSON(http.StatusOK, gin.H{
			"user": userTemp,
		})
	}
}
