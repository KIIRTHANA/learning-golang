// Topic - Printing and formatting strings

package main

import "fmt"

func main() {

	num := 15
	name := "Kirthana"

	// Print
	fmt.Print("HElloo ") // Does not add a new line to the end
	fmt.Print("World \n")

	// Println
	fmt.Println("Hey there!")
	fmt.Println("My lucky number is :", num)

	//Printf - formatted strings %_ = format specifier
	fmt.Printf("My name is %v and my age is %v \n", name, num)           // %v - var
	fmt.Printf("My name is %q and my age is %q \n", name, num)           // %q - quotes
	fmt.Printf("type of name is %T and type of num is %T \n", name, num) // %T - type of var
	fmt.Printf("My score is %f. \n", 234.456)                            // %f - flost
	fmt.Printf("My score is %0.2f \n", 234.456)                          // decimal places round off

	// Sprintf - save formatted strings
	sentence := fmt.Sprintf("My name is %v :P", name)
	fmt.Print(sentence)
	sent2 := fmt.Sprintf("My name is %q :P", name)
	fmt.Print(sent2)

}
