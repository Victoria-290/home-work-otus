package repository

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
)

var (
	usersMu  sync.RWMutex
	tasksMu  sync.RWMutex
	tokensMu sync.RWMutex

	users  []*user.User
	tasks  []*task.Task
	tokens []*auth.Token
)

const (
	usersFile  = "data/users.json"
	tasksFile  = "data/tasks.json"
	tokensFile = "data/tokens.json"
)

func init() {
	_ = os.MkdirAll("data", os.ModePerm)
	loadUsers()
	loadTasks()
	loadTokens()
}

// StoreUser сохраняет пользователя в слайс и файл
func StoreUser(u *user.User) {
	usersMu.Lock()
	defer usersMu.Unlock()
	users = append(users, u)
	saveToFile(usersFile, users)
}

// StoreTask сохраняет задачу в слайс и файл
func StoreTask(t *task.Task) {
	tasksMu.Lock()
	defer tasksMu.Unlock()
	tasks = append(tasks, t)
	saveToFile(tasksFile, tasks)
}

// StoreToken сохраняет токен в слайс и файл
func StoreToken(tok *auth.Token) {
	tokensMu.Lock()
	defer tokensMu.Unlock()
	tokens = append(tokens, tok)
	saveToFile(tokensFile, tokens)
}

// GetUsers возвращает копию пользователей
func GetUsers() []*user.User {
	usersMu.RLock()
	defer usersMu.RUnlock()
	return append([]*user.User(nil), users...)
}

// GetTasks возвращает копию задач
func GetTasks() []*task.Task {
	tasksMu.RLock()
	defer tasksMu.RUnlock()
	return append([]*task.Task(nil), tasks...)
}

// GetTokens возвращает копию токенов
func GetTokens() []*auth.Token {
	tokensMu.RLock()
	defer tokensMu.RUnlock()
	return append([]*auth.Token(nil), tokens...)
}

// saveToFile сериализует структуру в JSON и сохраняет в файл
func saveToFile(filename string, data interface{}) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	_ = enc.Encode(data)
}

// loadUsers загружает пользователей из файла
func loadUsers() {
	loadFromFile(usersFile, &users)
}

// loadTasks загружает задачи из файла
func loadTasks() {
	loadFromFile(tasksFile, &tasks)
}

// loadTokens загружает токены из файла
func loadTokens() {
	loadFromFile(tokensFile, &tokens)
}

// loadFromFile читает JSON и десериализует в слайс
func loadFromFile(filename string, target interface{}) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	_ = json.Unmarshal(data, target)
}
