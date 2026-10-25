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
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])
	// дополнительно
	fmt.Printf("Первое вхождение: %d\n", findStatus(statuses, "done"))
	fmt.Printf("Процент выполненного: %f", completionPercent(statuses))
}

func findStatus(statuses []string, target string) int {
	for ind, status := range statuses {
		if status == target {
			return ind
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	all := 0.0
	completed := 0.0
	for _, status := range statuses {
		if status == "done" {
			completed += 1
		}
		all += 1
	}
	if completed != 0 {
		return 100 * completed / all
	}
	return 0
}
