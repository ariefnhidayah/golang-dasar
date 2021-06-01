package main

import "fmt"

// function
func main() {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)
	// fmt.Println(add(1, 2))

	luas, keliling := calculate2(10, 2)

	fmt.Println("Luas: ", luas)
	fmt.Println("Keliling: ", keliling)

}

// func add(a int, b int) int {
// 	return a + b
// }

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar Golang"
// 	return newSentence
// }

// return 2 value
// func calculate(panjang int, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)

// 	return luas, keliling
// }

// function dengan predefined return value
func calculate2(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}
