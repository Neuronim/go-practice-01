package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Павел"
	group := "Учебная-2"
	var completed int = 1
	var interest bool = true

	fmt.Println("Участник: ", name)
	fmt.Println("Группа: ", group)
	fmt.Println("Курс: ", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес с бэкенду: %t\n", interest)
}
