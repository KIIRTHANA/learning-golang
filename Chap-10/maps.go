// Data type - map
// similar to dicts in python (k,v) [keys can be of any data type ] and all keys of the same map should have a single data type.

package main

import "fmt"

func main() {
	menu := map[string]float64{ //map[key type] val type
		"soup":         99.2,
		"french fries": 192.2,
		"fried rice":   180.3,
		"salad":        140.73,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	//looping
	fmt.Println("------- Menu --------")
	for k, v := range menu {
		fmt.Printf("%v Rs.%v \n", k, v)
	}

	// key as integer
	fmt.Println("\n-------- Cricketers ----------\n")
	scores := map[int]string{
		98: "Bhuvi",
		89: "Shami",
		76: "Bumrah",
		87: "dhoni",
		46: "yuvraj",
	}
	scores[87] = "Virat"
	scores[46] = "Sakaria"
	fmt.Println(scores) //saves in ascending order
	for k, v := range scores {
		fmt.Printf("%v scored %v \n", v, k)
	}
}
