package main

import "fmt"

func main() {
	const course = "Виноделие"               // название курса
	var name string = "Влад"                 // имя
	var group string = "Подготовительная-08" // название группы
	var labs int = 28                        // количество работ
	backend := false                         // интерес к бэкэнду
	labs = 20                                // меняем кол-во работ
	name = "Олег"                            // меняем имя
	// с const мы бы не смогли переназначить значение переменной
	fmt.Printf("Участник: %s\n", name)
	fmt.Println("Группа:", group)
	fmt.Println("Курс:", course)
	fmt.Println("Завершено работ:", labs)
	fmt.Println("Интерес к бэкенду:", backend)
}
