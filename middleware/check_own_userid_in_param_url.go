package middleware

import (
	"common"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckOwnUseridInParamUrl(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// get params from url
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		/*
			//dependencies
			store := storage.NewSQLServerStorage(db)
			business := biz.CreateStorage(store)
			// check exist usesr
			//user := business.GetUserByUserID()
		*/
		// get token
		token := c.GetHeader("token")
		// parse to claim
		claims, err := common.PraseToken(token)
		if err != nil {
			panic(common.TOKEN_NOT_APPROPRIATE)
		}
		// check is owned user
		if claims.UserID == userid {
			panic(common.USER_IS_NOT_OWNED)
		}
		// next
		c.Next()
	}
}
