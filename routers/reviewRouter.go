package routers

import (
	"main/entities"
	"main/middleware"
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getReviewRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/review")
	{
		//middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
		// no need permission
		user.GET("/:reviewid",
			transport.GetReviewStorage(db))
		user.GET("/list",
			transport.ListReviewAllTransport(db))
		user.GET("/list/:productid",
			transport.ListReviewBaseOnProductIDTransport(db))
		// need permission
		user.PUT("/:userid/:reviewid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			transport.UpdateReviewTransport(db))
		user.POST("/:userid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			transport.CreateReviewTransport(db))
		user.DELETE("/:userid/:reviewid",
			middleware.CheckOwnUseridInParamUrl(db, entities.ROLE_ADMIN),
			transport.DeleteReviewTransport(db))
	}
}
