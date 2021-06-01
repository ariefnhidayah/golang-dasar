package main

import "fmt"

func main() {

	// array / collection
	// var languages [5]string
	// languages := [5]string{
	// 	"Ruby",
	// 	"Python",
	// 	"Go",
	// 	"Java",
	// 	"PHP",
	// }

	// languages := [...]string{
	// 	"Ruby",
	// 	"Python",
	// 	"Go",
	// 	"Java",
	// 	"PHP",
	// 	"C#",
	// }

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)
	// fmt.Println("")

	// // looping array (foreach array)
	// for index, lang := range languages {
	// 	fmt.Println("Index: ", index, " Languange: ", lang)
	// }

	// untuk membuat array kosong yang lengthnya dynamic
	// var gamingConsoles []string

	// // untuk push value array / append value array
	// gamingConsoles = append(gamingConsoles, "PS 4")
	// gamingConsoles = append(gamingConsoles, "Nintendo")
	// gamingConsoles = append(gamingConsoles, "PS 2")

	// fmt.Println(gamingConsoles)

	// // iterate / mengambil value dari array
	// for _, console := range gamingConsoles {
	// 	fmt.Println(console)
	// }

	// Map
	// init map
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["test"] = 100
	// myMap["test"] = 90
	// myMap["testing"] = 200

	// fmt.Println(myMap["test"])

	// map dengan yang sudah ada isinya
	// myMap := map[string]string{
	// 	"test":        "Ini Test",
	// 	"testing":     "",
	// 	"HAHHAAHAHHA": "Tidak ada",
	// }

	// myMap["testing"] = "HAHAHA"

	// fmt.Println(myMap["testing"])

	// // iterasi map
	// for _, value := range myMap {
	// 	fmt.Println(" value: ", value)
	// }

	// fmt.Println("========")

	// delete(myMap, "test") // delete map

	// for _, value := range myMap {
	// 	fmt.Println(" value: ", value)
	// }

	// cek apakah map ada valuenya atau tidak
	// _, isAvail := myMap["aaa"] // a => value, b => bool

	// fmt.Println(isAvail)

	// slice of map
	students := []map[string]string{
		{
			"name":  "Arief",
			"score": "A",
		},
		{
			"name":  "Nur",
			"score": "B",
		},
		{
			"name":  "Hidayah",
			"score": "A",
		},
	}

	for _, student := range students {
		fmt.Println("Nama: ", student["name"], " Score: ", student["score"])
	}

}
