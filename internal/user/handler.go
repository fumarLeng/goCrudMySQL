package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc Service }

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (handler *Handler) Register(routerGroup *gin.RouterGroup) {
	g := routerGroup.Group("/users")
	g.POST("", handler.create)
	g.GET("", handler.list)
	g.GET(":id", handler.get)
	g.PUT(":id", handler.update)
	g.DELETE(":id", handler.delete)
}

// create godoc
// @Summary 建立使用者
// @Description 新增一筆使用者資料
// @Tags users
// @Param user body User true "使用者資料"
// @Router /users [post]
func (handler *Handler) create(c *gin.Context) {
	var in User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := handler.svc.Create(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, out)
}

// list godoc
// @Summary 取得所有使用者
// @Tags users
// @Produce json
// @Router /users [get]
func (handler *Handler) list(c *gin.Context) {
	us, err := handler.svc.List(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, us)
}

// get godoc
// @Summary 取得單一使用者
// @Tags users
// @Produce json
// @Param id path int true "使用者 ID"
// @Router /users/{id} [get]
func (handler *Handler) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	u, err := handler.svc.Get(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// update godoc
// @Summary 更新使用者
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "使用者 ID"
// @Param user body User true "使用者資料"
// @Router /users/{id} [put]
func (handler *Handler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var in User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := handler.svc.Update(c.Request.Context(), uint(id), in)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// delete godoc
// @Summary 刪除使用者
// @Tags users
// @Produce json
// @Param id path int true "使用者 ID"
// @Router /users/{id} [delete]
func (handler *Handler) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := handler.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
