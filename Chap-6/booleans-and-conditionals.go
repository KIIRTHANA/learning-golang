// Topic - Boolean and conditional statments in golang

package main

import "fmt"

func main() {
	age := 15
	/*
		fmt.Println(age <= 30)
		fmt.Println(age >= 30)
		fmt.Println(age == 30)
		fmt.Println(age != 30)
	*/

	if age > 30 {
		fmt.Println("age is less than 30")
	} else if age > 40 {
		fmt.Println("less than 40")
	} else {
		fmt.Println("age is not greater than 30 or 40")
	}

	// if inside for loop
	names := []string{"Arun", "Bean", "chiku", "dalv"}

	for index, values := range names {
		if index <= 2 {
			fmt.Println(values)
			continue
		} else if index > 2 {
			fmt.Println("breaking at 3rd")
			break
		}
	}
}
