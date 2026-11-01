package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

// Возвращает "#ID Название [статус], приоритет N"
func (t *Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

// Изменяет статус исходной задачи на done
func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	// TODO: найдите ID=2, завершите и установите приоритет 3.
	found := findTaskByID(tasks, 2)
	if found != nil {
		found.Complete()
		found.SetPriority(3)
	}

	// TODO: посчитайте завершённые задачи.
	completed := 0
	for _, task := range tasks {
		if task.Status == "done" {
			completed++
		}
		fmt.Println(task.Summary())
	}
	fmt.Println("Задач завершено: ", completed)

	// Ищем задачу с ID #99
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
