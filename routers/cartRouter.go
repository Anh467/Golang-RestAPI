package routers

import (
	"middleware"
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getCartRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/cart")
	{
		user.GET("/:userid/:productid", middleware.CheckOwnUseridInParamUrl(db), transport.GetCartTransport(db))
		user.GET("/:userid", middleware.CheckOwnUseridInParamUrl(db), transport.ListCartTransport(db))
		user.DELETE("/:userid/:productid", middleware.CheckOwnUseridInParamUrl(db), transport.DeleteCartTransport(db))
		user.PUT("/:userid/:productid", middleware.CheckOwnUseridInParamUrl(db), transport.UpdateCartTransport(db))
		user.POST("/:userid/:productid", middleware.CheckOwnUseridInParamUrl(db), transport.CreateCartTransport(db))
	}
}
