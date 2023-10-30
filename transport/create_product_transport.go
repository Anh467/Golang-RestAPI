package transport

import (
	"biz"
	"encoding/json"
	"entities"
	"net/http"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProductTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare product
		var product *entities.Product
		// Read data from request body
		requestBody := c.MustGet("requestBody").(map[string]interface{})
		productMap := requestBody["product"].(map[string]interface{})
		// Chuyển map thành dữ liệu JSON
		productJSON, err := json.Marshal(productMap)
		if err != nil {
			// Xử lý lỗi nếu có
			panic(err)
		}
		// Unmarshal dữ liệu JSON vào biến product kiểu dữ liệu của bạn
		err = json.Unmarshal(productJSON, &product)
		if err != nil {
			// Xử lý lỗi nếu có
			panic(err)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// get Data
		productTemp := business.CreateProductBiz(c, product)
		// res
		c.JSON(http.StatusOK, gin.H{
			"product": productTemp,
		})
	}
}
