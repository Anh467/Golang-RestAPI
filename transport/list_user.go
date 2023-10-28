package transport

import (
	"biz"
	"encoding/json"
	"entities"
	"net/http"
	"storage"

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
		users, err := business.ListNewUserModel(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// query
		db.Select(
			"UserID",
			"FullName",
			"Email",
			"Role").Table(entities.UserModelTable).Find(&users)
		// reponse
		c.JSON(http.StatusOK, gin.H{
			"users":  users,
			"length": len(users),
		})

	}
}

func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user entities.UserCreateModel
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = json.Unmarshal(body, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
}
