package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juanpicasti/go-todo-app/config"
	"strings"
)

func GetCorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	allowedOriginsString := config.CFG.AllowedOrigins
	allowedOrigins := strings.Split(allowedOriginsString, ",")
	corsConfig.AllowOrigins = allowedOrigins
	return cors.New(corsConfig)
}
