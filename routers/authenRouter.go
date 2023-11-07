package routers

import (
	"main/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAuthenRouters(api *gin.RouterGroup, db *gorm.DB) {
	user := api.Group("/authen")
	{
		user.POST("/signin", transport.SignInAuthen(db))
	}
}
