package routers

import (
	"entities"
	"middleware"
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/user")
	// get user
	//user.Use(middleware.Recovery())
	user.GET("/list", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.ListUser(db))
	user.POST("/create", transport.CreateUser(db))

}
