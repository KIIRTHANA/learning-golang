package main

import (
	"fmt"
	"strings"
)

// returning two strings
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	// to check if we pass first name only
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	fmt.Println(getInitials("kirthana sub"))
	fmt.Println(getInitials("arun venkat"))
	fmt.Println(getInitials("ishan"))

	fn, ln := getInitials("s n madhan nayaka")
	fmt.Println(fn, ln)

}
