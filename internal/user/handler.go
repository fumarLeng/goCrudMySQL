package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"goCrudMySQL/common"
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
// @Param user body CreateUserRq true "使用者資料"
// @Router /users [post]
func (handler *Handler) create(c *gin.Context) {
	var in CreateUserRq
	if err := c.ShouldBindJSON(&in); err != nil {
		common.Fail(c, common.CodeInvalidParam, "請求參數格式錯誤："+err.Error())
		return
	}
	out, err := handler.svc.Create(c.Request.Context(), in)
	if err != nil {
		code, msg := common.MapError(err)
		common.Fail(c, code, msg)
		return
	}
	common.Success(c, out)
}

// list godoc
// @Summary 取得所有使用者
// @Tags users
// @Produce json
// @Router /users [get]
func (handler *Handler) list(c *gin.Context) {
	us, err := handler.svc.List(c.Request.Context())
	if err != nil {
		common.Fail(c, http.StatusInternalServerError, "取得使用者清單失敗："+err.Error())
		return
	}
	common.Success(c, us)
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
		common.Fail(c, http.StatusBadRequest, "使用者 ID 格式錯誤")
		return
	}
	u, err := handler.svc.Get(c.Request.Context(), uint(id))
	if err != nil {
		common.Fail(c, http.StatusNotFound, "找不到指定的使用者："+err.Error())
		return
	}
	common.Success(c, u)
}

// update godoc
// @Summary 更新使用者
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "使用者 ID"
// @Param user body UpdateUserRequest true "更新使用者資料"
// @Success 200 {object} UserResponse
// @Router /users/{id} [put]
func (handler *Handler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.Fail(c, http.StatusBadRequest, "使用者 ID 格式錯誤")
		return
	}
	var in UpdateUserRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		common.Fail(c, http.StatusBadRequest, "請求參數格式錯誤："+err.Error())
		return
	}
	u, err := handler.svc.Update(c.Request.Context(), uint(id), in)
	if err != nil {
		common.Fail(c, http.StatusNotFound, "更新失敗："+err.Error())
		return
	}
	common.Success(c, u)
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
		common.Fail(c, http.StatusBadRequest, "使用者 ID 格式錯誤")
		return
	}
	if err := handler.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		common.Fail(c, http.StatusNotFound, "刪除失敗："+err.Error())
		return
	}
	common.Success(c, nil)
}
