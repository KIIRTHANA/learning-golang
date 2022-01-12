// Topic - pass by value (Go is known as pass by value lang)
// when var and its value is passed into another function, only a copy of it is passed and stored.
// making a change inside that new func won't change the value of original var, unless you return the value

package main

import "fmt"

// Group A - array, structs, bool, string, integer, float
func updateOrder(s string) string {
	s = "Peri Peri Chicken"
	return s
}

// Group B = func, slices, maps
func updatePrice(n map[string]int) {
	n["lemon tea"] = 15
	n["coffee"] = 18

}
func main() {
	food := "Fried Chicken"
	food = updateOrder(food)

	fmt.Println("The order is for:", food)

	items := map[string]int{
		"coffee": 20,
		"tea":    10,
		"juice":  30,
	}

	updatePrice(items) //new item is added from function
	fmt.Println(items) //updation also happens

	// Note - for group B elements the var stores the pointer to the values stored inside the var
	// when a var is sent to a func, it copies the pointer value!!

}
