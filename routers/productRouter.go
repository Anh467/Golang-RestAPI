package routers

import (
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getProductRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/product")
	{
		user.GET("/list", transport.ListProduct(db))
	}
}
