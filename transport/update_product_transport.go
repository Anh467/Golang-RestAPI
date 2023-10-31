package transport

import (
	"biz"
	"entities"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProductTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get data request body
		var product *entities.ProductUpdate
		if err := c.ShouldBindJSON(&product); err != nil {
			panic(err)
		}
		// get data param url
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(err)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.CreateStorage(store)
		// update
		productTemp := bussiness.UpdateProduct(c, *product, productid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"product": productTemp,
		})
	}
}
