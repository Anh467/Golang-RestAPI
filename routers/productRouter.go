package routers

import (
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getProductRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/product")
	{
		user.GET("/list/:productid", transport.GetProduct(db))
		user.GET("/list", transport.ListProduct(db))
		user.POST("/list")
		user.PUT("/list/:productid")
		user.DELETE("/list/:productid")
	}
}
