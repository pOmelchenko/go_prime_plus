package main

import "fmt"

// Упражнение по программированию 2.8
//
// В языке С одна функция может вызывать другую. Напишите программу, которая вызывает функцию по имени
// <code>one_three()</code>. Эта функция должна вывести слово "один" в одной строке, вызвать функцию <code>two()</code>,
// а затем вывести слово "три" тоже в одной строке. Функция <code>two()</code> должна отобразить слово "два" в одной
// строке. Функция <code>main()</code> должна вывести слово "начинаем:" перед вызовом функции <code>one_three()</code>
// и слово "порядок!" после ее вызова. Таким образом, выходные данные должны иметь следующий вид:
// начинаем: один
// два
// три
// порядок!
func main() {
	fmt.Print("начинаем: ")
	one_three()
	fmt.Println("порядок!")

}

func one_three() {
	fmt.Println("один")
	two()
	fmt.Println("три")
}

func two() {
	fmt.Println("два")
}
