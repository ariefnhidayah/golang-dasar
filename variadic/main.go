package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{1,5,6,1,2,6}
	avg := calculate(numbers...)
	fmt.Println("Rata-rata: ", avg)

	hobbies := []string{"Play Game", "Swim", "Learn Programming"}

	yourHobbies("Arief", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0

	for _, number := range(numbers) {
		total += number
	}

	avg := float64(total) / float64(len(numbers))

	return avg
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Println("My Name: ", name)
	fmt.Println("My Hobbies: ", hobbiesAsString)
}