package routers

import (
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
		user.PUT("/:reviewid",
			middleware.GetTokenInformation(db),
			transport.UpdateReviewTransport(db))
		user.POST("/",
			middleware.GetTokenInformation(db),
			transport.CreateReviewTransport(db))
		user.DELETE("/:reviewid",
			middleware.GetTokenInformation(db),
			transport.DeleteReviewTransport(db))
	}
}
