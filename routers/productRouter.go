package routers

import (
	"main/entities"
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getProductRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/product")
	{
		user.GET("/get/:productid", transport.GetProduct(db))
		user.GET("/list", transport.ListProduct(db))
		user.POST("/", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.CreateProductTransport(db))
		user.PUT("/:productid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.UpdateProductTransport(db))
		user.DELETE("/:productid", middleware.CheckRole(db, entities.ROLE_ADMIN), transport.DeleteProductTransport(db))
	}
}
