package middleware

import (
	"go-test/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("APP_SECRET")
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Missing or malformed token",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		rawTokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		rawToken, err := jwt.Parse(rawTokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil || !rawToken.Valid {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid or expired token",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := rawToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid token claims",
			})
			c.AbortWithStatus((http.StatusUnauthorized))
			return
		}

		userId, ok := claims["userId"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid userId in token",
			})
		}

		c.Set("userId", int(userId))
		c.Next()
	}
}
