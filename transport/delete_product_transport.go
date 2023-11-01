package transport

import (
	"biz"
	"net/http"
	"storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProductTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get data param url
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(err)
		}
		// dependeciese
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.CreateStorage(store)
		// delete
		bussiness.DeleteProduct(c, productid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "delete done",
		})
	}
}