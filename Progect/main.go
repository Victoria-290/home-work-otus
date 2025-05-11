package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	// Создание контекста с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Каналы для обмена данными и завершения
	eventCh := make(chan service.EntityEvent, 10)
	done := make(chan struct{})

	// Запуск горутин
	go service.StartGenerator(ctx, eventCh)
	go service.StoreFromChannel(ctx, eventCh)
	go service.StartLogger(ctx, done)

	// Обработка сигнала завершения приложения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh // Ожидание сигнала

	log.Println("Shutdown signal received")

	// Отмена контекста и ожидание завершения логгера
	cancel()
	time.Sleep(500 * time.Millisecond) // Подождать завершения задач
	close(done)

	log.Println("Application gracefully stopped")
}
