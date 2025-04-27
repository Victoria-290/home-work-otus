package main

import (

	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	// Создаем канал для передачи сущностей
	entityChan := make(chan repository.EntityEvent)

	// Запускаем генератор сущностей
	go service.GenerateAndSendEntities(entityChan)

	// Запускаем обработчик сущностей в репозитории
	go repository.StoreFromChannel(entityChan)

	// Запускаем логгер для отслеживания изменений в слайсах
	go repository.StartLogger()

	// Блокируем main, чтобы программа не завершалась сразу
	select {}
}
