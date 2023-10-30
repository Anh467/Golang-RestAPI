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

func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user *entities.UserCreateModel
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
		// setup dependencies
		store := storage.NewSQLServerStorage(db)
		business := biz.NewCreateBiz(store)
		//
		users, err := business.CreateUserBiz(c.Request.Context(), user)
		if err != nil {
			panic(err.Error())
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			//return
		}
		// reponse
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
