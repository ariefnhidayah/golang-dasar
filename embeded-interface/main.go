package main

import (
	"fmt"
	"math"
)

// Interface bisa di-embed ke interface lain. Cara penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin di-embed ke dalam interface tujuan.

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Kubus struct {
	Sisi float64
}

func (kubus *Kubus) Volume() float64 {
	return math.Pow(kubus.Sisi, 3)
}

func (kubus *Kubus) Luas() float64 {
	return math.Pow(kubus.Sisi, 2) * 6
}

func (kubus *Kubus) Keliling() float64 {
	return kubus.Sisi * 12
}

func main() {
	var bangunRuang Hitung = &Kubus{4}

	fmt.Println("==KUBUS==")
	fmt.Println("Volume: ", bangunRuang.Volume())
	fmt.Println("Luas: ", bangunRuang.Luas())
	fmt.Println("Keliling: ", bangunRuang.Keliling())
}
