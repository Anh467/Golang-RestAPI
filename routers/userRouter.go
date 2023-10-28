package routers

import (
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getListUser(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/user")
	{
		user.GET("/list", transport.ListUser(db))

	}
}
