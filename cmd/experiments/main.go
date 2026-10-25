package main

import "fmt"

func main() {
	const language = "Go"
	var name string = "Павел"
	completed := 2.367889
	var ready bool = true
	var hours float64 = 1.4675

	fmt.Println("Изучаем", language)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено: %.4f\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность: %.4f часа\n", hours)

	completed = 2
	fmt.Printf("После изменения: %.4f\n", completed)
}
