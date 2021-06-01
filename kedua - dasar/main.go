package main

import "fmt"

// variable
func main() {

	// // menggunakan var
	// var name string = "Arief"

	// // tidak menggunakan var
	// age := 20
	// fmt.Println(name)
	// fmt.Println(age)

	// percabangan
	// if
	// if else
	// else if
	// if bersarang

	// score := 80
	// var grade string

	// if score <= 50 {
	// 	grade = "E"
	// } else if score <= 60 {
	// 	grade = "C"
	// } else if score <= 68 {
	// 	grade = "D"
	// } else if score <= 79 {
	// 	grade = "C"
	// } else if score <= 88 {
	// 	grade = "B"
	// } else {
	// 	if score >= 90 {
	// 		grade = "A+"
	// 	} else {
	// 		grade = "A"
	// 	}
	// }
	// fmt.Println(grade)

	// switch
	// number := 10

	// switch number {
	// case 2:
	// 	fmt.Println("Dua")
	// case 1:
	// 	fmt.Println("Satu")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("Gaada")
	// }

	// for
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Looping :", i)
	// }

	// while (gaya go)
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Looping: ", i)
	// 	i++
	// }

	// loop foreach
	title := "Golang the best language in the world!"

	for _, letter := range title {
		fmt.Println("Letter: ", string(letter))
	}

}
