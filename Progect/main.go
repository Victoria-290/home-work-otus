package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	// Создаем контекст с отменой, чтобы управлять завершением работы всех горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для передачи событий (EntityEvent) от генератора к обработчику
	events := make(chan service.EntityEvent, 100)

	// Запуск генератора данных: пользователей, задач, токенов
	go service.StartGenerator(ctx, events)

	// Запуск асинхронного обработчика событий — сохранение в память и файлы
	go service.StoreFromChannel(ctx, events)

	// Запуск логгера, который отслеживает появление новых структур
	go service.StartLogger(ctx)

	// Канал для прослушивания сигналов ОС (SIGINT, SIGTERM)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Блокируемся до получения сигнала завершения
	<-sig
	cancel() // Отправляем сигнал завершения всем горутинам через контекст

	// После cancel() все горутины должны корректно завершиться
}
