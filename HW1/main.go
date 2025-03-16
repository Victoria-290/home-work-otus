package main

import (
	"fmt"
)

func main() {
	fmt.Println("Привет новый курс в Otus!")

	// Ввод чисел
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	// Вызов функции сложения
	sum := add(a, b)
	fmt.Printf("Сумма чисел %d и %d равна %d\n", a, b, sum)
}

// Простая функция сложения двух чисел
func add(x int, y int) int {
	return x + y
}
