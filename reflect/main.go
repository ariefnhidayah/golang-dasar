package main

import (
	"fmt"
	"reflect"
)

// Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya. Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.


func main() {
	number := 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("Tipe variable: ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai Variable: ", reflectValue.Int())
	}

	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)
}