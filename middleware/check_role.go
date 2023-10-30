package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(ctx context.Context, role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer recover
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": r,
				})
			}

		}()
		c.Next()
	}
}
