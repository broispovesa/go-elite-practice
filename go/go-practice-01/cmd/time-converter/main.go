package main

import "fmt"

func main() {
	var totalminutes int
	fmt.Print("Введите количество страниц: ")
	_, err := fmt.Scan(&totalminutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	if totalminutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным")
		return
	}
	const mins = 60
	fmt.Printf("Полных часов: %d\n", totalminutes/mins)
	fmt.Printf("Оставшихся минут: %d\n", totalminutes%mins)
}
