package main

import "fmt"

func sayHello(n string) {
	fmt.Printf("Good morning %v from Bangalore \n", n)
}

//function with single argument
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

//function with mutliple arguments
func printDetails(name string, regnum int, gpa float64) {
	fmt.Printf("\nThe student %v register num %v has secured %v. \n", name, regnum, gpa)
}

//passing a type slice of strings and a func as an argument

func studNames(n []string, f func(string)) {
	//iterate through each string inside n and call func for each string
	for _, v := range n {
		f(v) //for each item in n, f is called
	}
}

// return a value from function
func totalCost(prices []float64) float64 {
	total := 0.0
	for _, v := range prices {
		total += v
	}
	return total
}

func main() {

	sayHello("Varun")
	sayBye("Akshara")

	printDetails("Kiru Natsie", 18792, 3.8)

	fmt.Println(" ")
	studNames([]string{"Annie", "Alice", "Alana"}, sayHello)

	fmt.Println(" ")
	tc := totalCost([]float64{23.870, 20.9822, 30.098})
	fmt.Printf("The total amount is : %0.2f", tc)
}
