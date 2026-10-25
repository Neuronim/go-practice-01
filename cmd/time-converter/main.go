package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	// Этот блок ввода дан готовым; вычисления добавьте самостоятельно.
	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const minutesInHour int = 60
	var fullHours int = totalMinutes / minutesInHour
	var ostMinutes = totalMinutes % minutesInHour

	fmt.Printf("Полных часов: %d\n", fullHours)
	fmt.Printf("Оставшихся минут: %d\n", ostMinutes)
}
