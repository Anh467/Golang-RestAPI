package routers

import (
	"entities"
	"middleware"
	"transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getOrderRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/order")
	{
		user.POST("/:userid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			middleware.CheckRole(db, entities.ROLE_ADMIN),
			transport.CreateOrderTransport(db))
		user.PUT("/:userid/:orderid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			middleware.CheckRole(db, entities.ROLE_ADMIN),
			transport.UpdateOrderTransport(db))
		user.DELETE("/:userid/:orderid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			middleware.CheckRole(db, entities.ROLE_ADMIN),
			transport.DeleteOrderTransport(db))
		user.GET("/:userid/:orderid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			middleware.CheckRole(db, entities.ROLE_ADMIN),
			transport.GetOrderTransport(db))
		user.GET("/:userid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			middleware.CheckRole(db, entities.ROLE_ADMIN),
			transport.ListOrderTransport(db))
	}
}
