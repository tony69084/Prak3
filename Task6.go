package main

import "fmt"

func main() {

	fmt.Println("Движение лифта.")
	floors := 24
	currentFloor := 1
	maxPass := 2
	currentCountPass := 0

	up := true
	avFloor1 := 4
	avFloor2 := 7
	avFloor3 := 10

	for {
		if currentFloor == 1 {
			up = true
			currentCountPass = 0
			fmt.Println("Прибытие на первый этаж, высадка пассажиров...")
		} else if currentFloor == floors {
			up = false
			fmt.Println("Прибытие на последний этаж")
		}
		if up {
			currentFloor++
		} else {
			currentFloor--
		}

		if !up && avFloor1 == currentFloor && currentCountPass < maxPass {
			currentCountPass++
			avFloor1 = 0
			fmt.Println("Вход пассажира. Текущий этаж:", currentFloor, "\tКол-во пассажиров:", currentCountPass)
		} else if !up && avFloor2 == currentFloor && currentCountPass < maxPass {
			currentCountPass++
			avFloor2 = 0
			fmt.Println("Вход пассажира. Текущий этаж:", currentFloor, "\tКол-во пассажиров:", currentCountPass)
		} else if !up && avFloor3 == currentFloor && currentCountPass < maxPass {
			currentCountPass++
			avFloor3 = 0
			fmt.Println("Вход пассажира. Текущий этаж:", currentFloor, "\tКол-во пассажиров:", currentCountPass)
		}

		if avFloor1 == 0 && avFloor2 == 0 && avFloor3 == 0 && up {
			break
		}
		fmt.Println("Этаж -", currentFloor)
	}
}
