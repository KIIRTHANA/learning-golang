package main

import "fmt"

var score = 99.8

func main() {

	//from greeting.go we get access to func "sayhi" and points
	//cause they're all inside same package "main"

	sayHi("Marley")

	for _, v := range points {
		fmt.Println(v)
	}
	// shows error, but works
	// run as - go run main.go greeetings.go

	showScore()

}
