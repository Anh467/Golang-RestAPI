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
	user.GET("/get", transport.GetUser(db))
	user.GET("/:userid", transport.GetUser(db))
	user.PUT("/:userid", transport.UpdateUserTransport(db))
	user.DELETE("/:userid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.DeleteUserTransport(db))
}
