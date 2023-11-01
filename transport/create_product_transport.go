package transport

import (
	"biz"
	"entities"
	"net/http"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProductTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare product
		var product *entities.ProductModel
		// Read data from request body
		/*
			requestBody := c.MustGet("requestBody").(map[string]interface{})
		*/
		// get request body
		/*
				var requestBody map[string]interface{}
				decoder := json.NewDecoder(c.Request.Body)
				decoder.Decode(&requestBody)
				// parse to map
				productMap := requestBody["product"].(map[string]interface{})
				// parse map to object
				productJSON, err := json.Marshal(productMap)
				if err != nil {
					panic(err)
				}

			// Unmarshal JSON data into product object
			err = json.Unmarshal(productJSON, &product)
			if err != nil {
				panic(err)
			}
		*/
		if err := c.ShouldBindJSON(&product); err != nil {
			panic(err.Error())
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
