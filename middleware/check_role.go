package middleware

import (
	"main/biz"
	"main/common"
	"main/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckRole(db *gorm.DB, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// flag to check panic
		flag := false
		// Read data from request body
		/*
			var requestBody map[string]interface{}
			decoder := json.NewDecoder(c.Request.Body)
			decoder.Decode(&requestBody)
		*/
		// Get key data "token"
		token := c.GetHeader("token")
		// parse token
		claims, err := common.PraseToken(token)
		// check claims
		if err != nil {
			panic(common.ROLE_USER_DENIED)
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
		// set flag for the next middleware use
		c.Set("flag", flag)
		// go to next
		if !flag {
			panic(common.ROLE_USER_DENIED)
		}
		//c.Set("requestBody", requestBody)
		c.Next()
	}
}
