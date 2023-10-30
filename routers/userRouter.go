package routers

import (
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/user")
	// get user
	user.GET("/list", transport.ListUser(db))
	user.POST("/create", transport.CreateUser(db))

}
