package user

import "time"

// User 使用者資料結構
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id" example:"1"`                                                                       // 使用者 ID
	Name      string    `gorm:"type:varchar(100);not null" json:"name" binding:"required" example:"Leo"`                                // 使用者名稱
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" binding:"required,email" example:"leo@example.com"` // 使用者 Email
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt" example:"2025-05-13T12:00:00Z"`                                         // 建立時間
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt" example:"2025-05-13T12:30:00Z"`                                         // 最後更新時間
}

// CreateUserRq 建立使用者的RQ
type CreateUserRq struct {
	Name  string `json:"name" binding:"required" example:"Leo"`                    // 使用者名稱
	Email string `json:"email" binding:"required,email" example:"leo@example.com"` // 使用者 Email
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" example:"Leo"` // 選填，使用 pointer 區分有無填寫
	Email *string `json:"email,omitempty" binding:"omitempty,email" example:"leo@example.com"`
}

type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"Leo"`
	Email     string    `json:"email" example:"leo@example.com"`
	CreatedAt time.Time `json:"createdAt" example:"2025-05-13T12:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-05-13T12:30:00Z"`
}
