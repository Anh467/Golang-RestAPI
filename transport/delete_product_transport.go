package transport

import (
	"main/biz"
	"main/storage"
	"net/http"
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
		bussiness := biz.NewCreateBiz(store)
		// delete
		bussiness.DeleteProductBiz(c, productid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"status": "delete done",
		})
	}
}
