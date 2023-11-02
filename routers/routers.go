package routers

import (
	"middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func V1Router(r *gin.Engine, db *gorm.DB) {
	r.Use(middleware.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "pinghhhhhhhhhhhhh",
		})
	})

	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			// user
			getUserRouters(api, db)
			// product
			getProductRouters(api, db)
			// authen
			getAuthenRouters(api, db)
			// category
			getCategoryRouters(api, db)
			// cart
			getCartRouters(api, db)
			// order
			getOrderRouters(api, db)
		}
	}
}
