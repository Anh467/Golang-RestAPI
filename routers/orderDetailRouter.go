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
		user.POST("/", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.CreateOrderDetailTransport(db))
		user.PUT("/", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.UpdateOrderDetailTransport(db))
		user.DELETE("/", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.DeleteOrderDetailTransport(db))
		user.GET("/", middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN), transport.ListOrderDetailTransport(db))
	}
}
