package main

import (
	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	eventChannel := make(chan service.EntityEvent, 100)

	// Запускаем асинхронную обработку сущностей
	go service.StoreFromChannel(eventChannel)

	// Запускаем логгер
	go service.StartLogger()

	// Запускаем генератор сущностей
	go service.StartGenerator(eventChannel)

	// Блокируем main чтобы не завершился
	select {}
}
