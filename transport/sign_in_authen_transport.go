package transport

import (
	"main/biz"
	"main/entities"
	"main/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignInAuthen(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user *entities.UserModel
		if err := c.ShouldBindJSON(&user); err != nil {
			panic(err.Error())
		}
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//
		userJwt, err := business.SignInAuthenBiz(c.Request.Context(), user.Email, user.Password)
		if err != nil {
			panic(err.Error())
		}
		// reponse
		c.JSON(http.StatusOK, gin.H{
			"user": userJwt,
		})
	}
}
