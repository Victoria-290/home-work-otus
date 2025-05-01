package service

import (
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
)

// EntityEvent — обертка для любой сущности, которую нужно сохранить
type EntityEvent struct {
	Entity any
}

// StoreFromChannel — слушает канал и сохраняет полученные сущности через репозиторий
func StoreFromChannel(ch <-chan EntityEvent) {
	for event := range ch {
		repository.Store(event.Entity)
	}
}
