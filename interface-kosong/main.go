package main

import(
	"fmt"
	"strings"
)

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

	// casting variabel interface kosong
	// Variabel bertipe interface{} bisa ditampilkan ke layar sebagai string dengan memanfaatkan fungsi print, seperti fmt.Println(). Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk string dari nilai aslinya.

	secret = 2
	// var number = secret * 10 // jika seperti ini akan error
	var number = secret.(int) * 10
	fmt.Println("Multiplied by 10: ", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favourite fruits")

	// casting interface kosong ke object pointer
	secret = &Person{Name: "Arief", Age: 20}
	name := secret.(*Person).Name
	fmt.Println(name)
}

type Person struct {
	Name string
	Age int
}