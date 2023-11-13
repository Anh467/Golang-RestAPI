package routers

import (
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getCartRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/cart")
	{
		user.GET("/:productid", middleware.GetTokenInformation(db), transport.GetCartTransport(db))
		user.GET("/", middleware.GetTokenInformation(db), transport.ListCartTransport(db))
		user.DELETE("/:productid", middleware.GetTokenInformation(db), transport.DeleteCartTransport(db))
		user.PUT("/:productid", middleware.GetTokenInformation(db), transport.UpdateCartTransport(db))
		user.POST("/:productid", middleware.GetTokenInformation(db), transport.CreateCartTransport(db))
	}
}
