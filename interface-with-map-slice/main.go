package main

import "fmt"

func main() {

	var persons = []map[string]interface{}{
		{
			"name": "Arief",
			"age": 20,
		},
		{
			"name": "Nur",
			"age": 19,
		},
		{
			"name": "Hidayah",
			"age": 18,
		},
	}

	for _, person := range persons {
		fmt.Println(person["name"], "age is", person["age"])
	}

	var randoms = []interface{}{
		map[string]interface{}{"name": "Arief", "age": 22},
		[]string{"Informatika", "Sistem Informasi"},
		"Arief",
	}

	for _, random := range randoms {
		fmt.Println(random)
	}

}