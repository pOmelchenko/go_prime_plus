package main

import "fmt"

// Упражнение по программированию 3.2
//
// Напишите программу, которая приглашает ввести некоторое значение в коде ASCII, например, 66, а затем
// выводит символ, которому соответствует введенный код.
func main() {
	var symbol int
	fmt.Print("Введите код символа: ")
	fmt.Scan(&symbol)
	fmt.Printf("Был введен символ: %c\n", symbol)
}
