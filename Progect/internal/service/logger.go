package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// StartLogger — периодически логирует новые добавленные сущности. Завершается по context.
func StartLogger(ctx context.Context, done <-chan struct{}) {
	var lastUsersCount, lastTasksCount, lastTokensCount int

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-done:
			return
		case <-ticker.C:
			users := repository.GetUsers()
			if len(users) > lastUsersCount {
				for _, u := range users[lastUsersCount:] {
					fmt.Printf("[Logger] New user added: ID=%d, Email=%s\n", u.ID, u.Email)
				}
				lastUsersCount = len(users)
			}

			tasks := repository.GetTasks()
			if len(tasks) > lastTasksCount {
				for _, t := range tasks[lastTasksCount:] {
					fmt.Printf("[Logger] New task added: ID=%d, Title=%s, OwnerID=%d\n", t.ID, t.Title, t.OwnerID)
				}
				lastTasksCount = len(tasks)
			}

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
