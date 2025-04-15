package auth

import "time"

// Token — структура access/refresh токена
type Token struct {
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	IsRefresh bool      `json:"is_refresh"`
}
