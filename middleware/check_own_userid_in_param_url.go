package middleware

import (
	"main/biz"
	"main/common"
	"main/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// param roles for exception that when roles passing wwho has this role will can do something
func CheckOwnUseridInParamUrl(db *gorm.DB, roles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// check admin
		// flag to check rol
		flag := false
		// declare role
		// Get key data "token"
		token := c.GetHeader("token")
		// parse token
		claims, err := common.PraseToken(token)
		// check claims
		if err != nil {
			panic(common.TOKEN_NOT_APPROPRIATE)
		}
		// get role
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//
		user := business.GetUserBiz(c.Request.Context(), claims.Email, claims.Password)
		// check role
		for _, role := range roles {
			if role == user.Role {
				flag = true
			}
		}
		// set flag
		c.Set("flag", flag)
		// go to next
		if flag {
			c.Next()
			return
		}

		// check own
		// get params from url
		userid, err := strconv.Atoi(c.Param("userid"))
		if err != nil {
			panic(common.ERR_INTEGER_WRONG_FORMAT)
		}
		// check is owned user
		if claims.UserID != userid {
			panic(common.USER_IS_NOT_OWNED)
		}
		// next
		c.Next()
	}
}
