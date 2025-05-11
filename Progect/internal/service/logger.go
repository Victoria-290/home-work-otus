package service

import (
	"fmt"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// StartLogger запускается в горутине и каждые 200 мс проверяет изменения в слайсах
// При появлении новых элементов — логирует их в консоль
func StartLogger() {
	var lastUsersCount, lastTasksCount, lastTokensCount int

	for {
		time.Sleep(200 * time.Millisecond)

		// Получаем текущее состояние слайсов из репозитория
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
