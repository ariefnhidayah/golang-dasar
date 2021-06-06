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

	fmt.Println("=====")

	var student = &Student{
		Name:  "Arief Nur Hidayah",
		Grade: 2,
	}
	student.getPropertyInfo()

	fmt.Println("======")

	fmt.Println("Nama: ", student.Name)
	reflectValue = reflect.ValueOf(student)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("Nama: ", student.Name)
}

// Reflect bisa digunakan untuk mengambil informasi semua property variabel objek cetakan struct, dengan catatan property-property tersebut bermodifier public. Langsung saja kita praktekan, siapkan sebuah struct bernama student.

type Student struct {
	Name  string
	Grade int
}

func (student *Student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(student)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama: ", reflectType.Field(i).Name)
		fmt.Println("Tipe Data: ", reflectType.Field(i).Type)
		fmt.Println("Nilai: ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (student *Student) SetName(name string) {
	student.Name = name
}
