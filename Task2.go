package main

import "fmt"

func main() {

	var num1, num2 int
	i := num1
	fmt.Println("Сумма двух чисел по единице.")
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)

	for i != num1+num2 {
		i++
	}
	fmt.Println("Сумма двух чисел по единце равна:", i)
}
