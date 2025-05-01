package main

import (
	"github.com/Victoria-290/home-work-otus/Progect/internal/repository"
	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	// Канал для передачи сущностей между генератором и хранилищем
	eventChan := make(chan service.EntityEvent)

	// Запуск горутины генератора (создает сущности и отправляет в канал)
	go service.StartGenerator(eventChan)

	// Запуск горутины асинхронного сохранения (слой сервиса, не репозиторий)
	go service.StoreFromChannel(eventChan)

	// Запуск горутины логгера (сервис отслеживает изменения в хранилище)
	go service.StartLogger()

	// Блокировка main, чтобы горутины не завершились сразу
	select {}
}
