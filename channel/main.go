package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	// var sayHello = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHello("Arief")
	// go sayHello("Nur")
	// go sayHello("Hidayah")

	// var message1 = <-messages
	// fmt.Println(message1)

	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)

	texts := []string{"Arief", "Nur", "Hidayah"}

	for _, text := range texts {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(text)
	}

	for i := 0; i < len(texts); i++ {
		PrintMessage(messages)
	}
}

// send variabel channel melalui parameter func
func PrintMessage(what chan string) {
	fmt.Println(<-what)
}