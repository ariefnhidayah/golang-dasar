package main

import "fmt"

func main() {

	// Closure adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.

	// closure disimpan sebagai variabel
	var getMinMax = func(numbers []int) (int, int) {
		var min, max int

		for index, number := range numbers {
			switch {
			case index == 0:
				max, min = number, number
			case number > max:
				max = number
			case number < min:
				min = number
			}
		}

		return min, max
	}

	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var min, max = getMinMax(numbers)

	fmt.Println("Data: ", numbers)
	fmt.Println("Max: ", max)
	fmt.Println("Min: ", min)

	// Immediately-Invoked Function Expression (IIFE)
	// Closure jenis ini dieksekusi langsung pada saat deklarasinya.
	fmt.Println("=================================")
	min = 3
	var newNumbers = func(min int) []int {
		var temp []int
		for _, number := range numbers {
			if number < min {
				continue
			}
			temp = append(temp, number)
		}

		return temp
	}(min)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	// Closure sebagai nilai kembalian
	fmt.Println("=================================")

	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	max = 3
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany) // 9
	fmt.Println("value \t:", theNumbers)

}

// Closure sebagai nilai kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var temp []int

	for _, number := range numbers {
		if number <= max {
			temp = append(temp, number)
		}
	}

	return len(temp), func() []int {
		return temp
	}
}
