package transport

import (
	"main/biz"
	"main/common"
	"main/entities"
	"main/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateReviewTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var review entities.ReviewCreate
		// get userid
		userid := c.MustGet("userid").(int)
		// get body
		if err := c.ShouldBind(&review); err != nil {
			panic(common.JSON_BODY_WRONG_FORMAMT)
		}
		review.UserID = userid
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// update
		reviewTemp := bussiness.CreateReviewBiz(c, review)
		// res
		c.JSON(http.StatusOK, gin.H{
			"review": reviewTemp,
		})
	}
}
