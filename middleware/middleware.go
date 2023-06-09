package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/helpers"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
