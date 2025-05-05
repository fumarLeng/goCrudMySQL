package router

import (
	"github.com/gin-gonic/gin"
	"goCrudMySQL/internal/user"
)

func New(r *gin.Engine, uh *user.Handler) *gin.Engine {
	v1 := r.Group("/api/v1")
	uh.Register(v1)
	return r
}
