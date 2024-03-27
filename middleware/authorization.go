package middleware

import (
	"mygramapi/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result gin.H
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			result = gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, result)
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
