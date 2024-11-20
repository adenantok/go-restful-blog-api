package middleware

import (
	"go-restful-blog-api/v2/auth/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenuser := c.GetHeader("Authorization")
		if tokenuser == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Validasi token di sini
		claims, err := token.ValidateToken(tokenuser) // Menggunakan error untuk menentukan validitas token
		if err != nil {                               // Jika ada error, token invalid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan userID di context untuk digunakan di controller
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
