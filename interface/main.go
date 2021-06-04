package main

import (
	"fmt"
	"math"
)

// Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.
type BangunDatar interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	Diameter float64
}

func (lingkaran Lingkaran) JariJari() float64 {
	return lingkaran.Diameter / 2
}

func (lingkaran Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(lingkaran.JariJari(), 2)
}

func (lingkaran Lingkaran) Keliling() float64 {
	return math.Pi * lingkaran.Diameter
}

type Persegi struct {
	Sisi float64
}

func (persegi Persegi) Luas() float64 {
	return math.Pow(persegi.Sisi, 2)
}

func (persegi Persegi) Keliling() float64 {
	return persegi.Sisi * 4
}

func main() {
	var bangunDatar BangunDatar

	bangunDatar = Persegi{10.0} // sisi persegi 10
	fmt.Println("==PERSEGI==")
	fmt.Println("Luas: ", bangunDatar.Luas())
	fmt.Println("Keliling: ", bangunDatar.Keliling())

	fmt.Println("")

	bangunDatar = Lingkaran{14.0}
	fmt.Println("==LINGKARAN==")
	fmt.Println("Luas: ", bangunDatar.Luas())
	fmt.Println("Keliling: ", bangunDatar.Keliling())
	fmt.Println("Jari-Jari: ", bangunDatar.(Lingkaran).JariJari())
}
