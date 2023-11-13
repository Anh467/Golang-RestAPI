package middleware

import (
	"main/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// param roles for exception that when roles passing wwho has this role will can do something
func GetTokenInformation(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare role
		// Get key data "token"
		token := c.GetHeader("token")
		// parse token
		claims, err := common.PraseToken(token)
		// check claims
		if err != nil {
			panic(common.TOKEN_NOT_APPROPRIATE)
		}
		// set key
		c.Set("userid", claims.UserID)
		// next
		c.Next()
	}
}
