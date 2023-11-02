package transport

import (
	"biz"
	"net/http"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProduct(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		productid := c.Param("productid")
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// get product
		product := bussiness.GetProductBiz(c, productid)
		c.JSON(http.StatusOK, gin.H{
			"product": product,
		})
	}
}
