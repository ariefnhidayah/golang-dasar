package main

import "fmt"

func main() {

	var secret interface{}

	secret = "arief hidayah"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// penerapan interface dengan map
	data := map[string]interface{}{
		"name":    "Arief",
		"umur":    20,
		"hobbies": []string{"Playing game"},
	}

	fmt.Println(data["name"], data["umur"])
}
