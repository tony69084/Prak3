package main

import (
	"fmt"
)

func main() {

	// Емкость 3х корзин
	var appleCount1, appleCount2, appleCount3 int
	// Заполненность 3х корзин
	var basketCount1, basketCount2, basketCount3 int

	fmt.Println("Задача на постепенное наполнение корзин разной ёмкости (if и continue и break).")
	fmt.Println("Введите емкость первой корзины:")
	fmt.Scan(&appleCount1)

	fmt.Println("Введите емкость второй корзины:")
	fmt.Scan(&appleCount2)

	fmt.Println("Введите емкость третьей корзины:")
	fmt.Scan(&appleCount3)

	for i := 0; true; i++ {
		fmt.Println("Яблок в первой корзине:", basketCount1)
		fmt.Println("Яблок во второй корзине:", basketCount2)
		fmt.Println("Яблок в третьей корзине:", basketCount3, "\n")

		if i%3 == 0 {
			if basketCount1 < appleCount1 {
				basketCount1++
				continue
			} else {
				i++
			}
		}
		if i%3 == 1 {
			if basketCount2 < appleCount2 {
				basketCount2++
				continue
			} else {
				i++
			}
		}
		if i%3 == 2 {
			if basketCount3 < appleCount3 {
				basketCount3++
				continue
			} else {
				i++
			}
		}
		if basketCount1+basketCount2+basketCount3 == appleCount1+appleCount2+appleCount3 {
			break
		}
	}
	fmt.Println("Все корзины заполнены.")
}
