package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helper "github.com/teddylethal/golang-jwt-project/helpers"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, errorMsg := helper.ValidateToken(clientToken)
		if errorMsg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorMsg})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("user_id", claims.UserId)
		c.Set("user_role", claims.UserRole)
		c.Next()

	}
}
