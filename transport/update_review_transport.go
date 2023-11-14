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

func UpdateReviewTransport(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var review entities.ReviewUpdate
		// get userid
		userid := c.MustGet("userid").(int)
		// get headers
		reviewid, err := strconv.Atoi(c.Param("reviewid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// get body
		if err := c.ShouldBind(&review); err != nil {
			panic(common.JSON_BODY_WRONG_FORMAMT)
		}
		review.ReviewID = reviewid
		// dependencies
		store := storage.NewSQLServerStorage(db)
		bussiness := biz.NewCreateBiz(store)
		// update
		reviewTemp := bussiness.UpdateReviewBiz(c, review, userid)
		// res
		c.JSON(http.StatusOK, gin.H{
			"review": reviewTemp,
		})
	}
}
