package routers

import (
	"main/entities"
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getOrderDetailRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/orderdetail")
	{
		user.POST("/:userid/:orderid", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.CreateOrderDetailTransport(db))
		user.PUT("/:userid/:orderid/:productid", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.UpdateOrderDetailTransport(db))
		user.DELETE("/:userid/:orderid/:productid", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.DeleteOrderDetailTransport(db))
		user.GET("/:userid/:orderid", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.ListOrderDetailTransport(db))
	}
}
