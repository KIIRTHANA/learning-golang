// Topic - Loops

package main

import "fmt"

func main() {

	//x := 0
	// for loops in Go can be used as regular for or as 'while' loop as well

	/**
	for x < 5 {
		fmt.Println("value of x :", x)
		x++
	}
	**/

	/**
	for i := 0; i < 5; i++ {
		fmt.Println("value of x :", i)
	}
	**/

	names := []string{"Divya", "Chaitra", "Yashu", "Karishma"}

	/*
		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}
	*/

	// For in type of loop

	/*
		for index, value := range names {
			fmt.Println(index, value)
		} */

	// If you don't want to print index (vice-versa for value)
	for _, value := range names {
		fmt.Println(value)
	}

}
