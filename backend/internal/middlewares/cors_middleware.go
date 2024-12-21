package middlewares

import (
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/config"
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.GetEnv("FRONTEND_URL"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			log.Printf("Blocked CORS request from origin: %s", c.Request.Header.Get("Origin"))
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
