package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}
func (t *Task) SetPriority(stat int) bool {
	if stat < 1 || stat > 3 {
		return false
	}
	t.Priority = stat
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for t := range tasks {
		if tasks[t].ID == id {
			return &tasks[t]
		}
	}
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3}}
	ptr := findTaskByID(tasks, 2)
	ptr.Complete()
	ptr.SetPriority(3)
	k := 0
	for _, task := range tasks {
		fmt.Println(task.Summary())
		if task.Status == "done" {
			k += 1
		}
	}
	fmt.Printf("Задач завершено: %d\n", k)
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
