package main

import "fmt"

// pointer dengan struct
type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.Kom"
}

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// dengan var
	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(*numberB)
	// fmt.Println(numberB)

	// number := 5
	// fmt.Println("Alamat memory awal: ", &number)
	// fmt.Println("Nilai awal: ", number)

	// change(&number, 100)

	// fmt.Println("Alamat memory akhir: ", &number)
	// fmt.Println("Nilai akhir: ", number)

	// pointer struct
	student := Student{1, "Arief Nur Hidayah", 3.55}

	fmt.Println("Nama Awal: ", student.Name)

	student.graduate()

	fmt.Println("Nama Akhir: ", student.Name)
}

// func change(old *int, new int) {
// 	*old = new
// 	fmt.Println("Alamat memory function", old)
// 	fmt.Println("Didalam function: ", *old)
// }

// func graduate(student *Student) {
// 	student.Name = student.Name + " S.Kom"
// }
