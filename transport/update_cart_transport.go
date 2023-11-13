package transport

import (
	"main/biz"
	"main/common"
	"main/entities"
	"main/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCartTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get userid
		userid := c.MustGet("userid").(int)
		// declare
		var cart entities.CartUpdate
		// get param from header
		productid, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get param from request body
		if err := c.ShouldBindJSON(&cart); err != nil {
			panic(err)
		}
		cart.ProductID = productid
		cart.UserID = userid
		// dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		// creating
		cartTemp := business.UpdateCartBiz(c, cart)
		// res
		c.JSON(http.StatusOK, gin.H{
			"cart": cartTemp,
		})
	}
}
