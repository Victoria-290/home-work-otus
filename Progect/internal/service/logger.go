package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// StartLogger — горутина, которая логирует появление новых сущностей
func StartLogger(ctx context.Context) {
	var lastUsersCount, lastTasksCount, lastTokensCount int

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Logger] Завершение логгера...")
			return
		case <-ticker.C:
			// Проверка новых пользователей
			users := repository.GetUsers()
			if len(users) > lastUsersCount {
				for _, u := range users[lastUsersCount:] {
					fmt.Printf("[Logger] New user added: ID=%d, Email=%s\n", u.ID, u.Email)
				}
				lastUsersCount = len(users)
			}

			// Проверка новых задач
			tasks := repository.GetTasks()
			if len(tasks) > lastTasksCount {
				for _, t := range tasks[lastTasksCount:] {
					fmt.Printf("[Logger] New task added: ID=%d, Title=%s, OwnerID=%d\n", t.ID, t.Title, t.OwnerID)
				}
				lastTasksCount = len(tasks)
			}

			// Проверка новых токенов
			tokens := repository.GetTokens()
			if len(tokens) > lastTokensCount {
				for _, tok := range tokens[lastTokensCount:] {
					fmt.Printf("[Logger] New token added for UserID=%d, Expires=%s\n", tok.UserID, tok.ExpiresAt)
				}
				lastTokensCount = len(tokens)
			}
		}
	}
}
