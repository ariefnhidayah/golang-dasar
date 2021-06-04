package main

import (
	"fmt"
	"strings"
)

func main() {

	// contoh penggunaan
	data := []string{"Arief", "Nur", "Hidayah"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "i")
	})

	fmt.Println("Data asli: ", data)

	fmt.Println("Contains i: ", dataContainsO)

	dataLength5 := filter(data, func(each string) bool {
		return len(each) >= 6
	})

	fmt.Println("Data length >= 6: ", dataLength5)

}

// contoh function sebagai parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		filtered := callback(each)
		if filtered {
			result = append(result, each)
		}
	}

	return result
}
