package transport

import (
	"biz"
	"common"
	"encoding/json"
	"net/http"
	"storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Đọc dữ liệu từ request body
		var requestBody map[string]interface{}
		decoder := json.NewDecoder(c.Request.Body)
		decoder.Decode(&requestBody)

		// Lấy giá trị của trường "fullname"
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
		// get user
		user := business.GetUserBiz(c.Request.Context(), claims.Email, claims.Password)
		// res
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
