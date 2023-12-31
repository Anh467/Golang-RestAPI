package routers

import (
	"main/entities"
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getCategoryRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/category")
	{
		user.POST("/", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.CreateCategoryTransport(db))
		user.PUT("/:categoryid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.UpdateCategoryTransport(db))
		user.GET("/:categoryid", transport.GetCategoryTransport(db))
		user.GET("", transport.ListCategoryTransport(db))
		user.DELETE("/:categoryid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.DeleteCategoryTransport(db))
	}
}
