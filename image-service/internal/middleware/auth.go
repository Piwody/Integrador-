package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Acceso denegado: Token no proporcionado"})
			c.Abort()
			return
		}

		req, err := http.NewRequest("GET", "http://auth-service:8000/api/auth/verify", nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al verificar token"})
			c.Abort()
			return
		}
		req.Header.Add("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
