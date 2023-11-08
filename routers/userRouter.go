package routers

import (
	"main/entities"
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/user")
	// get user
	//user.Use(middleware.Recovery())
	user.GET("/list", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.ListUser(db))
	user.POST("/create", transport.CreateUser(db))
	// get the infomation of this user whose token
	user.GET("/get", transport.GetUser(db))
	// get the infomation of this user , but must have admin permission
	user.GET("/:userid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.GetUserByID(db))
	user.PUT("/:userid",
		middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
		transport.UpdateUserTransport(db))
	user.DELETE("/:userid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.DeleteUserTransport(db))
}
