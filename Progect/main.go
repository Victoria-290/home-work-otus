package main

import (
	"time"
	"github.com/Victoria-290/home-work-otus/Progect/internal/service"
)

func main() {
	// Запускаем генератор каждые 5 секунд
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			service.GenerateAndStoreEntities()
		}
	}
}
