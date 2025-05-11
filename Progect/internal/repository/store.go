package repository

import (
	"sync"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
)

var (
	usersMu  sync.Mutex
	tasksMu  sync.Mutex
	tokensMu sync.Mutex

	users  []*user.User
	tasks  []*task.Task
	tokens []*auth.Token
)

// StoreUser сохраняет пользователя потокобезопасно
func StoreUser(u *user.User) {
	usersMu.Lock()
	defer usersMu.Unlock()
	users = append(users, u)
}

// StoreTask сохраняет задачу потокобезопасно
func StoreTask(t *task.Task) {
	tasksMu.Lock()
	defer tasksMu.Unlock()
	tasks = append(tasks, t)
}

// StoreToken сохраняет токен потокобезопасно
func StoreToken(tok *auth.Token) {
	tokensMu.Lock()
	defer tokensMu.Unlock()
	tokens = append(tokens, tok)
}

// GetUsers возвращает копию списка пользователей
func GetUsers() []*user.User {
	usersMu.Lock()
	defer usersMu.Unlock()
	return append([]*user.User(nil), users...)
}

// GetTasks возвращает копию списка задач
func GetTasks() []*task.Task {
	tasksMu.Lock()
	defer tasksMu.Unlock()
	return append([]*task.Task(nil), tasks...)
}

// GetTokens возвращает копию списка токенов
func GetTokens() []*auth.Token {
	tokensMu.Lock()
	defer tokensMu.Unlock()
	return append([]*auth.Token(nil), tokens...)
}
