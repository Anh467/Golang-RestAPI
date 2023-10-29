package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func V1Router(r *gin.Engine, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "pinghhhhhhhhhhhhh",
		})
	})
	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			getUserRouters(api, db)
			getProductRouters(api, db)
		}
	}
}