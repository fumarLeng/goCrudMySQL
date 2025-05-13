package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goCrudMySQL/internal/user"
	"time"
)

func New(r *gin.Engine, uh *user.Handler) *gin.Engine {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // or ["http://localhost:8088"]
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1 := r.Group("/api/v1")
	uh.Register(v1)
	return r
}
