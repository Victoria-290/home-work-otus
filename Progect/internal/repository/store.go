package repository

import (
	"fmt"
	"sync"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
)

// Storable — универсальный интерфейс, который реализуют все сущности.
// Он позволяет передавать в Store() любые поддерживаемые типы.
type Storable interface {
	GetID() int64
}

// Мьютексы для каждого типа сущности —
// позволяют избежать блокировки всех операций при добавлении только одного типа данных.
var (
	usersMu  sync.Mutex
	tasksMu  sync.Mutex
	tokensMu sync.Mutex

	users  []*user.User
	tasks  []*task.Task
	tokens []*auth.Token
)

// Store сохраняет сущность в нужный слайс по типу
// Принимает сущность, реализующую интерфейс Storable,
// и добавляет её в соответствующее хранилище в памяти.
// Каждая ветка защищена своим мьютексом для минимизации блокировок.
func Store(s Storable) {
	switch v := s.(type) {

	// Добавление пользователя
	case *user.User:
		usersMu.Lock()
		defer usersMu.Unlock()
		users = append(users, v)
		fmt.Println("Stored User:", v.Email)

	// Добавление задачи
	case *task.Task:
		tasksMu.Lock()
		defer tasksMu.Unlock()
		tasks = append(tasks, v)
		fmt.Println("Stored Task:", v.Title)

	// Добавление токена
	case *auth.Token:
		tokensMu.Lock()
		defer tokensMu.Unlock()
		tokens = append(tokens, v)
		fmt.Println("Stored Token for user ID:", v.UserID)

	// Неизвестный тип — не сохраняем
	default:
		fmt.Println("Unknown type, not stored")
	}
}