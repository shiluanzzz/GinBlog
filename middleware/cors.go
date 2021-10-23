package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(
			cors.Options{
				//AllowAllOrigins:  true,
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
				AllowCredentials: true,
			},
		)
	}
}
