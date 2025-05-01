package service

import (
	"fmt"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// StartLogger запускает процесс логирования изменений в слайсах пользователей, задач и токенов.
// Проверка выполняется каждые 200 мс. Если найдены новые элементы — они выводятся в консоль.
func StartLogger() {
	var prevUsersLen, prevTasksLen, prevTokensLen int

	for {
		time.Sleep(200 * time.Millisecond)

		// Получаем копии текущих слайсов
		usersSnapshot := repository.GetUsersSnapshot()
		tasksSnapshot := repository.GetTasksSnapshot()
		tokensSnapshot := repository.GetTokensSnapshot()

		// Логируем новых пользователей
		if len(usersSnapshot) > prevUsersLen {
			newUsers := usersSnapshot[prevUsersLen:]
			for _, u := range newUsers {
				fmt.Printf("[Logger] New user added: ID=%d, Email=%s\n", u.ID, u.Email())
			}
			prevUsersLen = len(usersSnapshot)
		}

		// Логируем новые задачи
		if len(tasksSnapshot) > prevTasksLen {
			newTasks := tasksSnapshot[prevTasksLen:]
			for _, t := range newTasks {
				fmt.Printf("[Logger] New task added: ID=%d, Title=%s, OwnerID=%d\n", t.ID, t.Title(), t.OwnerID())
			}
			prevTasksLen = len(tasksSnapshot)
		}

		// Логируем новые токены
		if len(tokensSnapshot) > prevTokensLen {
			newTokens := tokensSnapshot[prevTokensLen:]
			for _, tok := range newTokens {
				fmt.Printf("[Logger] New token added: UserID=%d, ExpiresAt=%s\n", tok.UserID, tok.ExpiresAt.Format(time.RFC3339))
			}
			prevTokensLen = len(tokensSnapshot)
		}
	}
}
