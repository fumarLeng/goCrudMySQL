package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`           // 業務邏輯錯誤碼
	Message string      `json:"message"`        // 中文提示
	Data    interface{} `json:"data,omitempty"` // 實際資料
}

const (
	CodeOK             = 0
	CodeEmailExists    = 1001
	CodeUserNotFound   = 1002
	CodeInvalidParam   = 1003
	CodeInternalServer = 1999
)

// Success 成功回應
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeOK,
		Message: "success",
		Data:    data,
	})
}

// Fail 錯誤回應
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
