package repository

import (
	"sync"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
)

// Мьютексы для обеспечения конкурентного доступа к слайсам
var (
	usersMu  sync.RWMutex
	tasksMu  sync.RWMutex
	tokensMu sync.RWMutex
)

// Слайсы для хранения сущностей
var (
	users  []*user.User
	tasks  []*task.Task
	tokens []*auth.Token
)

// AddUser безопасно добавляет нового пользователя в слайс
func AddUser(u *user.User) {
	usersMu.Lock()
	defer usersMu.Unlock()
	users = append(users, u)
}

// AddTask безопасно добавляет новую задачу в слайс
func AddTask(t *task.Task) {
	tasksMu.Lock()
	defer tasksMu.Unlock()
	tasks = append(tasks, t)
}

// AddToken безопасно добавляет новый токен в слайс
func AddToken(tok *auth.Token) {
	tokensMu.Lock()
	defer tokensMu.Unlock()
	tokens = append(tokens, tok)
}

// GetUsersSnapshot возвращает копию текущего слайса пользователей
func GetUsersSnapshot() []*user.User {
	usersMu.RLock()
	defer usersMu.RUnlock()
	snapshot := make([]*user.User, len(users))
	copy(snapshot, users)
	return snapshot
}

// GetTasksSnapshot возвращает копию текущего слайса задач
func GetTasksSnapshot() []*task.Task {
	tasksMu.RLock()
	defer tasksMu.RUnlock()
	snapshot := make([]*task.Task, len(tasks))
	copy(snapshot, tasks)
	return snapshot
}

// GetTokensSnapshot возвращает копию текущего слайса токенов
func GetTokensSnapshot() []*auth.Token {
	tokensMu.RLock()
	defer tokensMu.RUnlock()
	snapshot := make([]*auth.Token, len(tokens))
	copy(snapshot, tokens)
	return snapshot
}
