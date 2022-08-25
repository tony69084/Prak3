package main

import (
	"fmt"
)

func main() {

	var productCount, discount float64

	fmt.Println("Расчёт суммы скидки.")
	fmt.Println("Введите цену товара:")
	fmt.Scan(&productCount)

	for {
		fmt.Println("Введите скидку (в процентах):")
		fmt.Scan(&discount)
		discountF := productCount * (discount / 100)
		if discount <= 30 && discountF <= 2000 {
			fmt.Println("Ваша скидка составит:", discountF)
			break
		}
		fmt.Println("Ваша скидка должна быть не больше 30% и не больше 2000 рублей, введите скидку снова")
	}
}
