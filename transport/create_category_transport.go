package transport

import (
	"biz"
	"entities"
	"net/http"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategoryTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare catgory
		var category entities.CategoryCreate
		// get data from request body
		if err := c.ShouldBindJSON(&category); err != nil {
			panic(err)
		}
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// creating
		categoryTemp := business.CreateCategoryBiz(c, category)
		// res
		c.JSON(http.StatusOK, gin.H{
			"category": categoryTemp,
		})
	}
}
