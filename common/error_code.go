package common

import "errors"

// 定義通用的錯誤
var (
	ErrEmailAlreadyExists = errors.New("email 已經被註冊")
	ErrInvalidUserID      = errors.New("使用者 ID 格式錯誤")
	ErrUserNotFound       = errors.New("找不到指定的使用者")
)
