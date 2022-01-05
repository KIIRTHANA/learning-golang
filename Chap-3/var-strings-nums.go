package main

import "fmt"

func main() {

	//Declaring variables three types
	var nameOne string = "Kira"
	var nameTwo = "Levi" //It'll automatically declare it as str
	var nameThree string

	nameThree = "Coni"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Eren"
	nameTwo = "Armin"

	// to declare a var - shortcut
	nameO := "Kiru"
	fmt.Println(nameOne, nameTwo, nameThree, nameO)

	//Integer Var declare and use
	var num1 int = 15
	var num2 = 19
	num3 := 25

	fmt.Println(num1, num2, num3)

	//bits and memory
	var num4 int8 = 15
	fmt.Println(num4)

	//float
	var num5 float32 = 23.4
	var num6 float64 = 87347234289.9
	num7 := 45.06 // by default it's assigned float64

	fmt.Println(num5, num6, num7)
}
