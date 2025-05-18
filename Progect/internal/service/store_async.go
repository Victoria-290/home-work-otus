package service

import (
	"context"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)
type EntityType int

const (
	UserType EntityType = iota
	TaskType
	TokenType
)

type EntityEvent struct {
	Type  EntityType
	Value interface{}
}

// StoreFromChannel — функция получения сущностей из канала и сохранения в соответствующий слайс, завершение по context
func StoreFromChannel(ctx context.Context, ch <-chan EntityEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-ch:
			switch e.Type {
			case UserType:
				if u, ok := e.Value.(*user.User); ok {
					repository.StoreUser(u)
				}
			case TaskType:
				if t, ok := e.Value.(*task.Task); ok {
					repository.StoreTask(t)
				}
			case TokenType:
				if tok, ok := e.Value.(*auth.Token); ok {
					repository.StoreToken(tok)
				}
			}
		}
	}
}