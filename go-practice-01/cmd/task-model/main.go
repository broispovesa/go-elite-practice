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

func (t Task) Complete() {
	t.Status = "done"
}

func main() {
	task := Task{ID: 1, Title: "Сделать README", Status: "new", Priority: 2}
	fmt.Println(task.Summary())
}
