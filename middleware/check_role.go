package middleware

import (
	"biz"
	"common"
	"encoding/json"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckRole(db *gorm.DB, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer recover
		flag := false
		// Đọc dữ liệu từ request body
		var requestBody map[string]interface{}
		decoder := json.NewDecoder(c.Request.Body)
		decoder.Decode(&requestBody)

		// Lấy giá trị của trường "fullname"
		token := requestBody["token"].(string)
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
				break
			}
		}

		if !flag {
			panic(common.ROLE_USER_DENIED)
		}

		c.Next()
	}
}