package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go Print(5, "halo")
	Print(5, "Apa kabar")

	var input string
	fmt.Scanln(&input)
}

func Print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
