package middleware

import (
	"assignment-final/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware(c *gin.Context)  {


        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wajib Menyertakan Authorization Header"})
			return
            
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

       
        claims, err := util.VerifyToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token Tidak Valid: %v", err)})
			return
            
        }
        c.Set("uuid", claims.UUID)
        c.Next()
}
