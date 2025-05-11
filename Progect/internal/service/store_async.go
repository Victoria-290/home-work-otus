package service

import (
	"context"

	"github.com/Victoria-290/home-work-otus/Progect/internal/model/auth"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/task"
	"github.com/Victoria-290/home-work-otus/Progect/internal/model/user"
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// StoreFromChannel — функция получения сущностей из канала и сохранения в соответствующий слайс, завершение по context
func StoreFromChannel(ctx context.Context, ch <-chan EntityEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-ch:
			switch v := e.(type) {
			case *user.User:
				repository.StoreUser(v)
			case *task.Task:
				repository.StoreTask(v)
			case *auth.Token:
				repository.StoreToken(v)
			}
		}
	}
}
