package transport

import (
	"main/biz"
	"main/entities"
	"main/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		//
		// declare	variable to store
		var users []entities.UserModel
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//
		users, err := business.ListNewUserModelBiz(c.Request.Context())
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// return
			panic(err.Error())
		}
		// reponse
		c.JSON(http.StatusOK, gin.H{
			"users":  users,
			"length": len(users),
		})

	}
}
