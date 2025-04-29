package router

import (
	"github.com/gin-gonic/gin"
	"goCrudMySQL/controller"
)

func NewRouter(u *controller.UserController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", u.Create)
			users.GET("", u.List)
			users.GET("/:id", u.Get)
			users.PUT("/:id", u.Update)
			users.DELETE("/:id", u.Delete)
		}
	}
	return r
}
