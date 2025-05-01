package repository

import (
	"log"
	"time"
)

// StartLogger - горутина для логирования изменений в слайсах
func StartLogger() {
	var prevUserCount, prevTaskCount, prevTokenCount int

	for {
		time.Sleep(200 * time.Millisecond)

		usersMu.Lock()
		currentUserCount := len(users)
		newUsers := users[prevUserCount:currentUserCount]
		usersMu.Unlock()

		tasksMu.Lock()
		currentTaskCount := len(tasks)
		newTasks := tasks[prevTaskCount:currentTaskCount]
		tasksMu.Unlock()

		tokensMu.Lock()
		currentTokenCount := len(tokens)
		newTokens := tokens[prevTokenCount:currentTokenCount]
		tokensMu.Unlock()

		if len(newUsers) > 0 {
			for _, u := range newUsers {
				log.Printf("New User Added: %s (ID %d)", u.Email, u.ID)
			}
			prevUserCount = currentUserCount
		}

		if len(newTasks) > 0 {
			for _, t := range newTasks {
				log.Printf("New Task Added: %s (ID %d)", t.Title, t.ID)
			}
			prevTaskCount = currentTaskCount
		}

		if len(newTokens) > 0 {
			for _, tk := range newTokens {
				log.Printf("New Token Added for user ID: %d", tk.UserID)
			}
			prevTokenCount = currentTokenCount
		}
	}
}
