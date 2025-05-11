package service

import (
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// EntityEvent — структура для передачи сущностей через канал
type EntityEvent struct {
	Type  string      // "user", "task", "token"
	Value interface{} // указатель на соответствующую структуру
}

// StoreFromChannel получает события из канала и вызывает соответствующие методы хранения
func StoreFromChannel(ch <-chan EntityEvent) {
	for e := range ch {
		switch e.Type {
		case "user":
			if u, ok := e.Value.(*user.User); ok {
				repository.StoreUser(u)
			}
		case "task":
			if t, ok := e.Value.(*task.Task); ok {
				repository.StoreTask(t)
			}
		case "token":
			if tok, ok := e.Value.(*auth.Token); ok {
				repository.StoreToken(tok)
			}
		default:
			// необработанный тип
		}
	}
}
