package user

import (
	"time"
)

// User — структура пользователя
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	password  string    // приватно — можно менять только через методы
	CreatedAt time.Time `json:"created_at"`
}

// NewUser создает нового пользователя
func NewUser(email, rawPassword string) *User {
	return &User{
		Email:     email,
		password:  rawPassword,
		CreatedAt: time.Now(),
	}
}

// CheckPassword сравнивает переданный пароль с внутренним
func (u *User) CheckPassword(input string) bool {
	// В реальности здесь будет hash сравнение
	return input == u.password
}

// SetPassword изменяет пароль (например, с хешированием)
func (u *User) SetPassword(newPassword string) {
	// Здесь можно захешировать, пока просто присваиваем
	u.password = newPassword
}
