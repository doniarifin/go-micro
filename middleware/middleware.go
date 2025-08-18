package middleware

import (
	helper "go-micro/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Invalid Authorization header"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helper.VerifyToken(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("exp", claims.RegisteredClaims.ExpiresAt)
		c.Set("userId", claims.RegisteredClaims.Issuer)

		claimsRole := claims.Role

		if role != "" && claimsRole != role {
			c.JSON(402, gin.H{"error": "forbidden access"})
			c.Abort()
		}

		c.Next()
	}
}
