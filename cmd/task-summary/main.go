package main

import "fmt"

// statusLabel превращает код статуса в слово
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество завершенных задач
func countStatus(statuses []string, target string) int {
	var counter int = 0
	for _, status := range statuses {
		if status == target {
			counter++
		}
	}
	return counter
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	// Добавляет еще один статус "new" в массив
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// Считает количество задач каждого статуса
	statusCounts := make(map[string]int)

	for _, status := range statuses {
		statusCounts[status]++
	}

	// Выводит статусы по порядку
	for _, status := range []string{"new", "in_progress", "done"} {
		fmt.Printf("%s: %d\n", statusLabel(status), statusCounts[status])
	}
}
