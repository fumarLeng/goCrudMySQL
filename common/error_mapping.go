package common

import (
	"errors"
)

// MapError 會對應 HTTP 狀態碼與訊息
func MapError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrEmailAlreadyExists):
		return CodeEmailExists, "建立使用者發生錯誤：Email 已經被註冊"
	case errors.Is(err, ErrUserNotFound):
		return CodeUserNotFound, "查無使用者"
	default:
		return CodeInternalServer, "系統內部發生錯誤"
	}
}
