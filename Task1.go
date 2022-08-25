package main

import "fmt"

func main() {

	var num int
	i := 0
	fmt.Println("Написание простого цикла.")
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	fmt.Println("Числа от 0 до", num)
	for i <= num {
		fmt.Println(i)
		i++
	}
}
