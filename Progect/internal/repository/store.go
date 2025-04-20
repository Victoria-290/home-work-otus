package repository

import (
	"fmt"
	"sync"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
	"github.com/Victoria-290/home-work-otus/Progect/internal/storable"
)

// Инициализируем "память"
var (
	users  []user.User
	tasks  []task.Task
	tokens []auth.Token

	mu sync.Mutex
)

// Store сохраняет сущность в нужный слайс по типу
func Store(s storable.Storable) {
	mu.Lock()
	defer mu.Unlock()

	switch v := s.(type) {
	case *user.User:
		users = append(users, *v)
		fmt.Println("Stored User:", v.Email)
	case *task.Task:
		tasks = append(tasks, *v)
		fmt.Println("Stored Task:", v.Title)
	case *auth.Token:
		tokens = append(tokens, *v)
		fmt.Println("Stored Token for user ID:", v.UserID)
	default:
		fmt.Println("Unknown type, not stored")
	}
}
