package service

import (
	"context"
	"strconv"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
)

// StartGenerator — горутина, создающая случайные структуры и отправляющая их в канал
func StartGenerator(ctx context.Context, ch chan<- EntityEvent) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var id int64 = 1

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Генерация пользователя
			u := &user.User{
				ID:    id,
				Email: "user" + strconv.FormatInt(id, 10) + "@example.com",
			}
			ch <- EntityEvent{Type: UserType, Value: u}

			// Генерация задачи
			t := &task.Task{
				ID:      id,
				Title:   "Task #" + strconv.FormatInt(id, 10),
				OwnerID: id,
			}
			ch <- EntityEvent{Type: TaskType, Value: t}

			// Генерация токена
			tok := &auth.Token{
				UserID:    id,
				Token:     "token" + strconv.FormatInt(id, 10),
				ExpiresAt: time.Now().Add(24 * time.Hour),
			}
			ch <- EntityEvent{Type: TokenType, Value: tok}

			id++
		}
	}
}
