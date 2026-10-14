package main

import "fmt"

func main() {
	first := 20
	second := 25
	third := 35
	var summ int = first + second + third
	var avarage float64 = float64(summ) / 3
	fmt.Printf("Сумма: %d\n Среднее: %.2f", summ, avarage)
}
