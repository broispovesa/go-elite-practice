package main

import "fmt"

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестно"
	}
}

func countStatus(statuses []string, target string) int {
	k := 0
	for _, element := range statuses {
		if element == target {
			k += 1
		}
	}
	return k
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")
	// statuses := []string{"blocked"}
	// fmt.Println(statusLabel("blocked"))
	fmt.Printf("Задач всего: %d\n", len(statuses))
	for ind, status := range statuses {
		fmt.Printf("%d. ", ind+1)
		fmt.Println(statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d", counts["new"], counts["in_progress"], counts["done"])
}
