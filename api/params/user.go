package params

import (
	"database/sql"
	"time"
)

type CreateUserRequest struct {
	UserName string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
}

type UpdateUserRequest struct {
	UserID   int64         `json:"user_id" binding:"required"`
	Email    string        `json:"email" binding:"required,email"`
	FullName string        `json:"full_name" binding:"required"`
	Gender   sql.NullInt16 `json:"gender"`
}

type UpdateUserPasswordRequest struct {
	UserID   int64  `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserResponse struct {
	UserID   int64         `json:"user_id" binding:"required"`
	Email    string        `json:"email" binding:"required,email"`
	FullName string        `json:"full_name" binding:"required"`
	Gender   sql.NullInt16 `json:"gender"`
}

type UserResponse struct {
	UserID                int64     `json:"user_id"`
	UserName              string    `json:"user_name" binding:"required,alphanum"`
	Email                 string    `json:"email" binding:"required,email"`
	AccessToken           string    `json:"access_token"`
	AccesstokenExpiredAt  time.Time `json:"access_token_expired_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
	CreatedAt             time.Time `json:"created_at"`
	ChangedPasswordAt     time.Time `json:"changed_password_at"`
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
}
